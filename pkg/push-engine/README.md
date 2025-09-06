# Git Push Engine — Quick Guide

A concise, day‑to‑day usage summary of the Git Push Engine. For the detailed manual (concepts, options, internals), see [ENGINE-MANUAL.md](util/git/ENGINE-MANUAL.md:1).

## What it does

- One command rolls up commits and pushes across the atelier/artist/canvas hierarchy.
- Correct order: canvases → artists → atelier.
- Single combined commit per level (when needed): stages working-tree changes and updated submodule pointers together.

Core entrypoint:
- Orchestrator: [push-engine.sh](util/git/push-engine.sh:1)
- Delegates: [canvas-push.sh](util/git/canvas-push.sh:1), [artist-push.sh](util/git/artist-push.sh:1), [atelier-push.sh](util/git/atelier-push.sh:1)
- Helpers and config: [git-helpers.sh](util/git/git-helpers.sh:1), [config.sh](util/git/config.sh:1)

## Quick start commands

Use the Makefile target at the level you are in:

- From a canvas directory:
  - make push
  - Rolls up and pushes this canvas only (no recursion below).

- From an artist directory:
  - make push
  - Recurses into this artist’s canvases, pushes canvases first, then makes a single combined artist commit (working tree + updated canvas pointers) and pushes it.

- From the atelier root:
  - make push
  - Recurses into all artists and canvases, pushes them first, then makes a single combined atelier commit (working tree + updated artist pointers) and pushes it.

That’s it. No flags required for recursion.

## Defaults that matter

- Auto-commit: ON by default
  - The engine stages working-tree changes and submodule pointer updates and creates a single commit per level when something is staged.
  - Default controlled by [config.sh](util/git/config.sh:36) via AUTO_COMMIT_DEFAULT=true.

- Non-interactive by default
  - Confirmations are auto-accepted by the orchestrator (see [push-engine.sh](util/git/push-engine.sh:14) and [bash.confirm_action()](util/git/git-helpers.sh:144)).

- Verbosity
  - LOG_LEVEL defaults to "info". For deeper diagnostics export LOG_LEVEL=debug when running make push.

## Typical day‑to‑day usage

- Normal commits:
  - Use standard git in the repo you’re working in (canvas, artist, or atelier).
  - Example: git add -A && git commit -m "feat: your change" && git push

- End‑of‑day roll‑up:
  - Run make push at the appropriate level:
    - Project-wide: at atelier root
    - One artist: in that artist directory
    - One canvas: in that canvas directory

The engine will detect and push what changed beneath that level and create a single combined commit in the current level if needed.

## Dry‑run preview

- To preview without changing anything:
  - util/git/push-engine.sh --dry-run
  - Works from any level; the engine detects the current level and prints actions. Dry‑run always exits 0 (see [push-engine.sh](util/git/push-engine.sh:133)).

## Overriding behavior (occasionally)

Environment variables (see [config.sh](util/git/config.sh:16)):

- AUTO_COMMIT_DEFAULT=false
  - Require manual staging; engine will only push already committed changes.

- LOG_LEVEL=debug
  - Print debug details from helpers and delegates (see [bash.log_debug()](util/git/git-helpers.sh:17)).

- Re‑enable interactive confirmations
  - ENGINE_ASSUME_YES=false and CONFIRM_PUSH_DEFAULT=true if you want to be prompted (default is non‑interactive).

Most users will not need overrides—make push is designed to work out‑of‑the‑box.

## Conventions and assumptions

- Structure and markers:
  - Atelier root contains .atelier; artists contain .artist; canvases contain .canvas (used by [bash.detect_level()](util/git/git-helpers.sh:25)).
- Naming:
  - Artists match "artist-*", canvases match "canvas-*" (see [config.sh](util/git/config.sh:40)).
- Remotes:
  - Each level uses an "origin" remote; new repos must have their initial remote set up once.

## More details

For full documentation—options, internals, discovery, failure isolation, exit codes, and troubleshooting—see the manual:
- [ENGINE-MANUAL.md](util/git/ENGINE-MANUAL.md:1)

Key scripts for reference:
- Orchestrator: [push-engine.sh](util/git/push-engine.sh:1)
- Canvas: [canvas-push.sh](util/git/canvas-push.sh:1)
- Artist: [artist-push.sh](util/git/artist-push.sh:1)
- Atelier: [atelier-push.sh](util/git/atelier-push.sh:1)
- Helpers: [git-helpers.sh](util/git/git-helpers.sh:1)
- Config: [config.sh](util/git/config.sh:1)