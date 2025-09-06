# Git Push Engine — Quick Guide

A concise, day‑to‑day usage summary of the Git Push Engine integrated with Atelier CLI. For the detailed manual (concepts, options, internals), see [ENGINE-MANUAL.md](pkg/push-engine/ENGINE-MANUAL.md:1).

## What it does

- One command rolls up commits and pushes across the atelier/artist/canvas hierarchy.
- Correct order: canvases → artists → atelier.
- Single combined commit per level (when needed): stages working-tree changes and updated submodule pointers together.

Core entrypoint:
- Orchestrator: [push-engine.sh](pkg/push-engine/push-engine.sh:1)
- Delegates: [canvas-push.sh](pkg/push-engine/canvas-push.sh:1), [artist-push.sh](pkg/push-engine/artist-push.sh:1), [atelier-push.sh](pkg/push-engine/atelier-push.sh:1)
- Helpers and config: [git-helpers.sh](pkg/push-engine/git-helpers.sh:1), [config.sh](pkg/push-engine/config.sh:1)

## Quick start commands

Use the Atelier CLI push commands from the appropriate directory level:

- From a canvas directory:
  - `atelier-cli canvas push`
  - Rolls up and pushes this canvas only (no recursion below).

- From an artist directory:
  - `atelier-cli artist push`
  - Recurses into this artist's canvases, pushes canvases first, then makes a single combined artist commit (working tree + updated canvas pointers) and pushes it.

- From the atelier root:
  - `atelier-cli push`
  - Recurses into all artists and canvases, pushes them first, then makes a single combined atelier commit (working tree + updated artist pointers) and pushes it.

That's it. No flags required for recursion.

## Defaults that matter

- Auto-commit: ON by default
  - The engine stages working-tree changes and submodule pointer updates and creates a single commit per level when something is staged.
  - Default controlled by [config.sh](pkg/push-engine/config.sh:36) via AUTO_COMMIT_DEFAULT=true.
  
  - Non-interactive by default
    - Confirmations are auto-accepted by the orchestrator (see [push-engine.sh](pkg/push-engine/push-engine.sh:14) and [bash.confirm_action()](pkg/push-engine/git-helpers.sh:144)).

- Verbosity
  - LOG_LEVEL defaults to "info". For deeper diagnostics export LOG_LEVEL=debug when running Atelier CLI push commands.

## Typical day‑to‑day usage

- Normal commits:
  - Use standard git in the repo you’re working in (canvas, artist, or atelier).
  - Example: git add -A && git commit -m "feat: your change" && git push

- End‑of‑day roll‑up:
  - Run Atelier CLI push commands from the appropriate level:
    - Project-wide: `atelier-cli push` from atelier root
    - One artist: `atelier-cli artist push` from artist directory
    - One canvas: `atelier-cli canvas push` from canvas directory

The engine will detect and push what changed beneath that level and create a single combined commit in the current level if needed.

## Dry‑run preview

- To preview without changing anything:
  - `atelier-cli push --dry-run` (from atelier root)
  - `atelier-cli artist push --dry-run` (from artist directory)
  - `atelier-cli canvas push --dry-run` (from canvas directory)
  - Works from any level; the engine detects the current level and prints actions. Dry‑run always exits 0.

## Overriding behavior (occasionally)

Environment variables (see [config.sh](pkg/push-engine/config.sh:16)):

- AUTO_COMMIT_DEFAULT=false
  - Require manual staging; engine will only push already committed changes.

- LOG_LEVEL=debug
  - Print debug details from helpers and delegates (see [bash.log_debug()](pkg/push-engine/git-helpers.sh:17)).

- Re‑enable interactive confirmations
  - ENGINE_ASSUME_YES=false and CONFIRM_PUSH_DEFAULT=true if you want to be prompted (default is non‑interactive).

Most users will not need overrides—make push is designed to work out‑of‑the‑box.

## Conventions and assumptions

- Structure and markers:
  - Atelier root contains .atelier; artists contain .artist; canvases contain .canvas (used by [bash.detect_level()](pkg/push-engine/git-helpers.sh:25)).
- Naming:
  - Artists match "artist-*", canvases match "canvas-*" (see [config.sh](pkg/push-engine/config.sh:40)).
- Remotes:
  - Each level uses an "origin" remote; new repos must have their initial remote set up once.

## More details

For full documentation—options, internals, discovery, failure isolation, exit codes, and troubleshooting—see the manual:
- [ENGINE-MANUAL.md](pkg/push-engine/ENGINE-MANUAL.md:1)

Key scripts for reference:
- Orchestrator: [push-engine.sh](pkg/push-engine/push-engine.sh:1)
- Canvas: [canvas-push.sh](pkg/push-engine/canvas-push.sh:1)
- Artist: [artist-push.sh](pkg/push-engine/artist-push.sh:1)
- Atelier: [atelier-push.sh](pkg/push-engine/atelier-push.sh:1)
- Helpers: [git-helpers.sh](pkg/push-engine/git-helpers.sh:1)
- Config: [config.sh](pkg/push-engine/config.sh:1)