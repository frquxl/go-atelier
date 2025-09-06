# Git Push Engine Manual

A detailed guide for using and understanding the Git Push Engine that orchestrates commits and pushes across the atelier/artist/canvas hierarchy.

## Quick Start

- From a canvas: run: make push
- From an artist: run: make push
- From the atelier root: run: make push

Defaults:
- Recurses automatically at artist and atelier levels (no flags required).
- Auto-commits staged and working-tree changes with a single combined commit per level.
- Non-interactive by default; confirmations are auto-accepted.

See the day-to-day summary in [README](util/git/README.md:1). For full details, continue below.

## Components

- Orchestrator: [push-engine.sh](util/git/push-engine.sh:1)
- Level delegates: [canvas-push.sh](util/git/canvas-push.sh:1), [artist-push.sh](util/git/artist-push.sh:1), [atelier-push.sh](util/git/atelier-push.sh:1)
- Helpers: [git-helpers.sh](util/git/git-helpers.sh:1)
- Config: [config.sh](util/git/config.sh:1)

Key helper functions:
- [bash.detect_level()](util/git/git-helpers.sh:25)
- [bash.find_artists()](util/git/git-helpers.sh:60), [bash.find_canvases()](util/git/git-helpers.sh:64)
- [bash.get_repo_info()](util/git/git-helpers.sh:38)
- [bash.has_uncommitted_changes()](util/git/git-helpers.sh:76), [bash.has_unpushed_commits()](util/git/git-helpers.sh:69)
- [bash.is_submodule_modified()](util/git/git-helpers.sh:86)
- [bash.safe_git_command()](util/git/git-helpers.sh:94)
- [bash.validate_git_repo()](util/git/git-helpers.sh:107), [bash.validate_remote()](util/git/git-helpers.sh:115)
- [bash.confirm_action()](util/git/git-helpers.sh:144)

## Usage

The engine detects the current level and delegates to the appropriate push script.

### Make targets

- Canvas directory: make push
  - Runs the engine scoped to the canvas. Non-recursive.
- Artist directory: make push
  - Recurses into all canvases, commits/pushes canvases first, then makes one combined artist commit and pushes.
- Atelier root: make push
  - Recurses into all artists and canvases, commits/pushes artists after canvases, then makes one combined atelier commit and pushes.

### Direct script usage

You can run the orchestrator directly:

bash
util/git/push-engine.sh [OPTIONS] [LEVEL_ARGS...]

Options:
- --dry-run: Preview actions; always exits 0 (see [push-engine.sh](util/git/push-engine.sh:133)).
- --quiet: Reduce verbosity.
- --force: Pass through to delegates for future use.

Level arguments (compatibility; not required with default recursion):
- Atelier: --artists, --all
- Artist: --canvases

Note: Default recursion at artist/atelier levels means you typically do not need these flags.

### Environment variables

From [config.sh](util/git/config.sh:16):
- DRY_RUN_DEFAULT=false
- VERBOSE_DEFAULT=true
- CONFIRM_PUSH_DEFAULT=true
- CONFIRM_FORCE_DEFAULT=false
- LOG_LEVEL_DEFAULT="info"
- AUTO_COMMIT_DEFAULT=true (engine default; override with AUTO_COMMIT_DEFAULT=false to require manual staging)
- AUTO_COMMIT_MESSAGE="engine: auto-commit uncommitted changes"

Additional behavior:
- Non-interactive confirmations are enabled by the orchestrator via ENGINE_ASSUME_YES=true in [push-engine.sh](util/git/push-engine.sh:14), which makes [bash.confirm_action()](util/git/git-helpers.sh:144) return success without prompting.
- You can force interactive prompts by exporting ENGINE_ASSUME_YES=false and CONFIRM_PUSH_DEFAULT=true.
- LOG_LEVEL=debug enables detailed [bash.log_debug()](util/git/git-helpers.sh:17) output.

### Typical workflows

- End-of-day roll-up at atelier: make push
- Ship all changes in a single artist: cd artist-name && make push
- Push a single canvas: cd artist-name/canvas-name && make push
- Dry-run preview anywhere: util/git/push-engine.sh --dry-run

## What the engine does

The engine ensures correct ordering and single-commit semantics at each level:

1) Canvas ([canvas-push.sh](util/git/canvas-push.sh:1))
- If there are uncommitted changes and AUTO_COMMIT_DEFAULT=true: stage and commit once; push.
- Otherwise, push any unpushed commits.

2) Artist ([artist-push.sh](util/git/artist-push.sh:1))
- Recurses into all canvases (delegate call).
- Stages updated canvas pointers and any artist working-tree changes.
- Creates one combined artist commit if anything is staged; pushes.

3) Atelier ([atelier-push.sh](util/git/atelier-push.sh:1))
- Recurses into all artists (each of which recurses into canvases).
- Stages updated artist pointers and any atelier working-tree changes.
- Creates one combined atelier commit if anything is staged; pushes.

### Detection and decisions

- Level detection: [bash.detect_level()](util/git/git-helpers.sh:25)
- Submodule modification: [bash.is_submodule_modified()](util/git/git-helpers.sh:86)
- Uncommitted changes: [bash.has_uncommitted_changes()](util/git/git-helpers.sh:76)
- Unpushed commits: [bash.has_unpushed_commits()](util/git/git-helpers.sh:69)
- Repo/remote validation: [bash.validate_git_repo()](util/git/git-helpers.sh:107), [bash.validate_remote()](util/git/git-helpers.sh:115)

### Commit strategy

- Stage working-tree changes (git add -A).
- Stage updated submodule pointers where detected.
- Make a single commit per level if the index is non-empty.
- Avoid empty commits by testing index via git diff --cached --quiet.

### Dry-run semantics

- Orchestrator always exits 0 in dry-run: [push-engine.sh](util/git/push-engine.sh:133).
- Delegates print planned actions but do not mutate state.
- Useful for previewing end-to-end effects before real pushes.

### Exit codes

Defined in [config.sh](util/git/config.sh:49):
- EXIT_SUCCESS=0
- EXIT_ERROR=1
- EXIT_INVALID_LEVEL=2
- EXIT_NO_CHANGES=3 (normalized to success by [push-engine.sh](util/git/push-engine.sh:139))
- EXIT_GIT_ERROR=4

### Failure isolation

- Delegation happens per item (artist/canvas) so a failure is logged but does not abort the complete atelier run.
- See per-item execution and warnings in [atelier-push.sh](util/git/atelier-push.sh:114) and [artist-push.sh](util/git/artist-push.sh:107).

## Initial setup and new repos

- The engine requires valid Git repos and an "origin" remote at each level.
- For a brand-new artist/canvas added by the CLI:
  - Configure the SSH remote and make an initial push once.
  - Thereafter, make push at any level will include the new repos automatically.

## Deletions and removals

- If an artist/canvas submodule is removed by the CLI and .gitmodules/gitlink is updated:
  - Parent level will detect changes, stage deletions, and create a commit that updates pointers.
  - Missing or invalid subrepos are logged and skipped; overall run continues.

## Troubleshooting

- Engine does nothing:
  - Ensure markers exist (.atelier/.artist/.canvas) and current directory is correct.
  - Ensure the repo has a configured "origin" remote.
- Engine prompts unexpectedly:
  - Export ENGINE_ASSUME_YES=false to re-enable prompts; otherwise it is non-interactive.
- Empty commit errors:
  - The engine avoids empty commits by checking the index; ensure you actually changed files.
- Pattern discovery misses repos:
  - Ensure artist-* and canvas-* naming is used.

## Reference

- Orchestrator: [push-engine.sh](util/git/push-engine.sh:1)
- Canvas delegate: [canvas-push.sh](util/git/canvas-push.sh:1)
- Artist delegate: [artist-push.sh](util/git/artist-push.sh:1)
- Atelier delegate: [atelier-push.sh](util/git/atelier-push.sh:1)
- Helpers: [git-helpers.sh](util/git/git-helpers.sh:1)
- Config: [config.sh](util/git/config.sh:1)