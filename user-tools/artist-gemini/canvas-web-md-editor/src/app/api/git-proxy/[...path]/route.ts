/* src/app/api/git-proxy/[...path]/route.ts
   A server-side proxy for isomorphic-git that injects Authorization for private GitHub repos
   to bypass browser CORS limitations and keep tokens off the client.
*/
import type { NextRequest } from 'next/server';

// Ensure Node runtime and dynamic route evaluation
export const runtime = 'nodejs';
export const dynamic = 'force-dynamic';

// Utility to (safely) build the upstream URL from catch-all segments + original search params
function buildUpstreamUrl(segments: string[], search: string): string {
  // Robustly reconstruct URL from catch-all segments used by isomorphic-git with corsProxy.
  // Examples we may receive:
  //  - ['https:', '', 'github.com', 'owner', 'repo.git', 'info', 'refs']
  //  - ['github.com', 'owner', 'repo.git', 'info', 'refs']  (without explicit protocol)
  if (!segments || segments.length === 0) {
    throw new Error('Missing upstream URL segments');
  }
  let s = segments.join('/');
  // Normalize: ensure protocol is present and properly formatted with double slash
  if (/^https?:\//i.test(s) && !/^https?:\/\//i.test(s)) {
    // Has "https:/" but not "https://"
    s = s.replace(/^https?:\//i, (m) => m + '/');
  }
  if (!/^https?:\/\//i.test(s)) {
    s = 'https://' + s;
  }
  return search ? `${s}${search}` : s;
}

function basicAuthHeaderFromToken(token: string): string {
  // GitHub supports token as username with 'x-oauth-basic' as password
  const encoded = Buffer.from(`${token}:x-oauth-basic`).toString('base64');
  return `Basic ${encoded}`;
}

// Common proxy handler
async function handleProxy(request: NextRequest, segments: string[]) {
  let upstreamUrl: string;
  try {
    upstreamUrl = buildUpstreamUrl(segments, request.nextUrl.search);
  } catch (e: any) {
    return new Response(`Proxy error: ${e?.message || 'Invalid upstream URL'}`, { status: 400 });
  }

  const isGithub = /:\/\/github\.com\//i.test(upstreamUrl) || /:\/\/api\.github\.com\//i.test(upstreamUrl);

  // Read server-side token (prefer private server env vars, fallback to NEXT_PUBLIC_ if needed)
  const token =
    process.env.GITHUB_TOKEN ||
    process.env.GH_TOKEN ||
    process.env.PERSONAL_ACCESS_TOKEN ||
    process.env.NEXT_PUBLIC_GITHUB_TOKEN;

  // Prepare headers: forward minimal required headers and inject Authorization for GitHub
  const forwardHeaders = new Headers();

  // Forward some client headers that are relevant for git smart HTTP
  const accept = request.headers.get('accept');
  const contentType = request.headers.get('content-type');
  // isomorphic-git sets "Git-Protocol" (case-insensitive)
  const gitProtocol = request.headers.get('git-protocol') || request.headers.get('Git-Protocol');

  if (accept) forwardHeaders.set('accept', accept);
  if (contentType) forwardHeaders.set('content-type', contentType);
  if (gitProtocol) forwardHeaders.set('Git-Protocol', gitProtocol);

  // Explicit UA (some providers check this)
  forwardHeaders.set('user-agent', 'isomorphic-git-proxy/1.0 (+nextjs)');

  // Inject Authorization for GitHub upstreams
  if (isGithub) {
    if (!token) {
      return new Response('GitHub token not configured on server (.env GITHUB_TOKEN)', { status: 401 });
    }
    forwardHeaders.set('authorization', basicAuthHeaderFromToken(token));
  }

  // Copy the request body if present (GET/HEAD have no body)
  const method = request.method.toUpperCase();
  let body: BodyInit | undefined = undefined;
  if (method !== 'GET' && method !== 'HEAD') {
    // Read body as ArrayBuffer to avoid streaming/duplex issues
    const arrayBuffer = await request.arrayBuffer();
    body = arrayBuffer;
  }

  // Perform the upstream request
  const upstreamResp = await fetch(upstreamUrl, {
    method,
    headers: forwardHeaders,
    body,
  });

  // Propagate upstream errors with body text for easier debugging
  if (!upstreamResp.ok) {
    const errText = await upstreamResp.text().catch(() => upstreamResp.statusText);
    return new Response(errText || upstreamResp.statusText, {
      status: upstreamResp.status,
      statusText: upstreamResp.statusText,
      headers: new Headers({ 'content-type': 'text/plain' }),
    });
  }

  // Build response back to client, streaming body
  const respHeaders = new Headers();
  const upstreamCT = upstreamResp.headers.get('content-type');
  if (upstreamCT) respHeaders.set('content-type', upstreamCT);
  const cc = upstreamResp.headers.get('cache-control');
  if (cc) respHeaders.set('cache-control', cc);

  return new Response(upstreamResp.body, {
    status: upstreamResp.status,
    statusText: upstreamResp.statusText,
    headers: respHeaders,
  });
}

// GET handler
export async function GET(request: NextRequest, context: { params: Promise<{ path: string[] }> }) {
  const { path } = await context.params;
  try {
    return await handleProxy(request, path);
  } catch (e: any) {
    return new Response(`Proxy error: ${e?.message || 'Unknown error'}`, { status: 500 });
  }
}

// POST handler (needed for git-upload-pack / git-receive-pack)
export async function POST(request: NextRequest, context: { params: Promise<{ path: string[] }> }) {
  const { path } = await context.params;
  try {
    return await handleProxy(request, path);
  } catch (e: any) {
    return new Response(`Proxy error: ${e?.message || 'Unknown error'}`, { status: 500 });
  }
}