# Cobra CLI Template Bootstrap Plan

## Overview

This document outlines the plan to implement a **Cobra CLI framework bootstrap** for the `go-atelier` project. When a user creates a new artist named `golang` with the `--cobra` flag, the CLI will automatically scaffold a Cobra CLI-first development environment with specialized context for AI agents and developers.

## Rationale

Currently, artist templates (sketch, gallery, default) provide generic boilerplate. This proposal adds a specialized template for developers working on CLI applications using the Cobra framework, enabling:

- Pre-configured Cobra project structure and best practices
- AI-aware context (AGENTS.md) explaining CLI-first development patterns
- Makefile targets optimized for CLI development
- Immediate productivity for CLI projects without manual setup

## Feasibility Assessment

✅ **Status: HIGHLY FEASIBLE**

The existing architecture already supports this pattern:

### Existing Template System
- **Template selection** already exists in `pkg/engine/engine.go` (lines 88-94) based on artist name
- **Embedded templates** use Go's `//go:embed` in `pkg/templates/templates.go` (line 11)
- **Multiple template types** already supported: `artist-default`, `artist-sketch`, `artist-gallery`
- **Boilerplate generation** is generic and works for any `projectType` parameter

### Precedent
- Special canvas handling for "sunflowers" (lines 178-184 in `engine.go`) demonstrates the pattern for injecting specialized assets
- Framework-specific templates can follow the same pattern

## Implementation Plan

### Phase 1: Create Template Assets

#### 1.1 Directory Structure
Create `pkg/templates/assets/artist-golang-cobra/` with the following files:

```
pkg/templates/assets/artist-golang-cobra/
├── README.md           # Cobra CLI introduction and project structure
├── AGENTS.md           # AI context for CLI-first development with Cobra
├── Makefile            # Cobra-specific build targets
├── gitignore           # Go + Cobra patterns
└── geminiignore        # AI filtering rules
```

#### 1.2 File Contents

**README.md** - Should cover:
- Welcome message for CLI-first developers
- Cobra framework basics and philosophy
- Typical file structure for a Cobra CLI project
- Quick-start commands (make scaffold, make build, make test)
- Link to Cobra documentation
- Context about single vs. multi-command CLIs

**AGENTS.md** - Should cover:
- AI context explaining that this artist specializes in CLI development
- Cobra best practices (command structure, flags, subcommands)
- Common patterns (global flags, persistent commands, hook pattern)
- Generator best practices (cobra-cli init)
- Testing patterns for CLI applications
- Guidance on avoiding common pitfalls
- Links to canonical Cobra examples

**Makefile** - Should include targets:
```makefile
.PHONY: scaffold build test clean fmt

scaffold:
	cobra-cli init
	# Generates a new Cobra CLI scaffold

build:
	go build -o bin/cli ./cmd/main.go

test:
	go test -v ./...

fmt:
	go fmt ./...
	gofmt -s -w .

clean:
	rm -rf bin/
	go clean
```

**gitignore** - Should include:
```
# Go
bin/
dist/
*.exe
*.dll
*.so

# IDE
.vscode/
.idea/
*.swp
*.swo

# Cobra CLI generated files
cmd/root.go
cmd/root_test.go
.cobra.yaml
```

**geminiignore** - Should prevent:
```
# Prevent AI from querying environment variables
env
.env*

# Prevent overwriting established patterns
Makefile
README.md
AGENTS.md
```

---

### Phase 2: Wire Into Artist Creation

#### 2.1 Modify `cmd/artist.go`

Add a `--cobra` flag to `artistInitCmd`:

```go
artistInitCmd.Flags().Bool("cobra", false, "Bootstrap artist with Cobra CLI framework template")
```

Extract the flag in the RunE function:

```go
withCobra, _ := cmd.Flags().GetBool("cobra")
```

#### 2.2 Modify `pkg/engine/engine.go`

Update the `CreateArtist()` function signature:

```go
func CreateArtist(atelierPath, artistName, canvasName string, frameworkTemplate string) (err error)
```

Add framework template selection logic in `CreateArtist()`:

```go
// Determine which template to use based on artistName or explicit framework
templateType := "artist-default"
if frameworkTemplate == "cobra" {
	templateType = "artist-golang-cobra"
} else if artistName == "sketch" {
	templateType = "artist-sketch"
} else if artistName == "gallery" {
	templateType = "artist-gallery"
}
```

#### 2.3 Update Call Sites

Update all calls to `CreateArtist()` in:
- `cmd/artist.go` - pass `--cobra` flag value
- `cmd/init.go` - pass empty string (use defaults)

---

### Phase 3: Optional - Interactive Prompt (Enhancement)

For better UX, add interactive framework selection when creating artists:

```go
// In artistInitCmd.RunE()
var framework string
if cmd.Flags().Changed("cobra") {
	framework = "cobra"
} else {
	// Optionally prompt for framework selection
	framework = promptFrameworkSelection() // could be empty for default
}
```

---

### Phase 4: Documentation Updates

#### 4.1 Update Main README.md

Add section under "Usage":

```markdown
### Add a New Artist with Cobra CLI Template

```bash
# Create an artist with Cobra CLI scaffolding
cd atelier-my-project
atelier-cli artist init golang --cobra
```

This creates an artist pre-configured for CLI-first development with the Cobra framework.
```

#### 4.2 Update IDEAS.md

Mark item #18 as:
```markdown
18. ✅ CLI first pattern - IMPLEMENTED: artist-golang can pre-set Cobra CLI framework with specialized AGENTS.md context
```

---

## Technical Specifications

### Backward Compatibility
- ✅ All changes are backward compatible
- Default behavior (no flags) remains unchanged
- Existing `--with-canvas` flag continues to work
- Existing artist types (sketch, gallery) unaffected

### File Structure Changes
```
pkg/templates/assets/
├── artist-default/          # Existing
├── artist-sketch/           # Existing
├── artist-gallery/          # Existing
├── artist-golang-cobra/     # NEW
│   ├── README.md
│   ├── AGENTS.md
│   ├── Makefile
│   ├── gitignore
│   └── geminiignore
└── canvas/                  # Existing
```

### Code Impact
- `cmd/artist.go`: +1 flag definition, +2 lines of extraction
- `cmd/init.go`: +1 parameter in CreateArtist calls
- `pkg/engine/engine.go`: +2-3 lines in template selection logic
- `pkg/templates/templates.go`: No changes needed (already generic)

---

## Implementation Checklist

### Phase 1: Template Creation
- [ ] Create `pkg/templates/assets/artist-golang-cobra/` directory
- [ ] Write `README.md` with Cobra introduction
- [ ] Write `AGENTS.md` with AI context for CLI development
- [ ] Write `Makefile` with CLI build targets
- [ ] Create `gitignore` with Go + Cobra patterns
- [ ] Create `geminiignore` with AI filtering

### Phase 2: Code Integration
- [ ] Add `--cobra` flag to `artistInitCmd` in `cmd/artist.go`
- [ ] Extract `--cobra` flag in RunE function
- [ ] Update `CreateArtist()` signature to accept framework parameter
- [ ] Add framework template selection logic in `CreateArtist()`
- [ ] Update all `CreateArtist()` call sites
- [ ] Test artist creation with and without `--cobra` flag

### Phase 3: Testing
- [ ] Write E2E test for `atelier-cli artist init golang --cobra`
- [ ] Verify template files are embedded correctly
- [ ] Verify AGENTS.md context is appropriate for AI agents
- [ ] Test with existing flags (`--with-canvas`)

### Phase 4: Documentation
- [ ] Update main `README.md` with Cobra artist example
- [ ] Update `IDEAS.md` to mark item #18 as complete
- [ ] Add entry to `AGENTS.md` about framework bootstrapping

### Phase 5: Future Enhancements (Optional)
- [ ] Add `--rust`, `--python` support for other CLI frameworks
- [ ] Add interactive framework selection prompt
- [ ] Create canvas-specific templates (e.g., `canvas-cobra-example`)

---

## Example Usage

### Create a Golang Cobra CLI Artist

```bash
$ cd atelier-my-project
$ atelier-cli artist init golang --cobra

Initializing artist...
[...]
Initializing canvas...
[...]
Connecting artist to atelier...
[...]
Artist 'golang' initialized successfully in atelier 'my-project'!

$ ls artist-golang/
README.md          # Contains Cobra introduction
AGENTS.md          # Contains AI context for CLI development
Makefile           # Contains CLI build targets
.gitignore
.geminiignore
.artist
canvas-example/
```

### Result

The artist is now pre-configured for CLI development:
- Developers understand they're in a Cobra CLI context
- AI agents receive clear context about Cobra patterns and best practices
- Makefile provides appropriate CLI development targets
- AGENTS.md prevents AI from overwriting established conventions

---

## Considerations

### Template Maintenance
- AGENTS.md should be kept current with Cobra best practices
- Consider referencing official Cobra documentation
- Update Makefile targets based on Go ecosystem changes

### Naming and Clarity
- Framework templates are artist-specific (not canvas-specific)
- Could expand to other languages in future (artist-rust-cobra, artist-python-click)
- Pattern is generalizable for other frameworks

### AI Awareness
- AGENTS.md is critical for guiding AI agents toward Cobra patterns
- Should include examples of good CLI structure
- Should discourage AI from generating non-idiomatic Cobra code

### Scope Control
- This MVP focuses only on Cobra + Go
- Other frameworks can be added independently following the same pattern
- Each template is self-contained and doesn't affect others

---

## Success Criteria

✅ When complete:
1. User can create an artist with `atelier-cli artist init golang --cobra`
2. Artist receives Cobra-specific templates and context
3. AGENTS.md properly guides AI agents on CLI best practices
4. Makefile has CLI-appropriate build targets
5. All existing functionality remains unchanged
6. Documentation clearly explains the feature

---

## Future Enhancements

### Post-MVP Ideas
1. **Framework Selection Prompt** - Interactive menu for framework choice
2. **Multi-Language Support** - Templates for Rust CLI, Python Click, Node.js Commander
3. **Canvas-Specific Templates** - Pre-scaffold a Cobra command structure in canvas
4. **Version Pinning** - Store preferred Cobra/Go versions in artist context
5. **CI/CD Integration** - Pre-configure GitHub Actions for CLI testing

---

## References

- **Cobra Documentation**: https://cobra.dev
- **Existing Template System**: `pkg/templates/templates.go`
- **Artist Creation Logic**: `pkg/engine/engine.go` (lines 62-136)
- **Related IDEAS**: Item #18 in `IDEAS.md`
- **Precedent**: Special handling for "sunflowers" canvas (lines 178-184 in `engine.go`)
