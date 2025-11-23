# Go Compiler Fork: race2 Detector (ARCHIVED)

> **⚠️ This repository is archived and no longer maintained.**  
> **New development has moved to: https://github.com/amirkhaki/moriarty**

---

## What This Was

An experimental fork of the Go compiler (Go 1.26-devel) that added a `-race2` flag by copying the existing race detector infrastructure and replacing `race` with `race2` everywhere.

**The approach**: Find all places where the original `-race` detector hooks in, duplicate the code, and change the names. Simple and hacky.

**Why it's archived**: Forking the entire Go compiler is unmaintainable. The upstream moves too fast, merge conflicts pile up, and it's overkill for what we actually need.

---

## What We Did

Basically grep'd for `race` throughout the codebase and added `race2` variants:

**Compiler** (`src/cmd/compile/`):
- Added `-race2` flag next to `-race` flag
- Created `IsRaceEnabled()` helper that returns `race || race2`
- Created `RacePrefix()` helper that returns `"race"` or `"race2"`
- Changed instrumentation code to use these helpers instead of hardcoding `race`
- Added `race2*` functions to `builtin.go` (regenerate with the build script)

**Runtime** (`src/runtime/`):
- Created `race2.go` with `//go:build race2` - copies of all race functions
- Created `race20.go` with `//go:build !race2` - no-op stubs
- Everywhere the runtime calls `race*` functions, added calls to `race2*` too
- Added `race2init()` call in `schedinit()`

**Build system** (`src/cmd/go/`):
- Added `BuildRace2` flag
- Copy-pasted race detector platform checks for race2
- Made build tags work with `race2`

**Files touched**: ~50 files (mostly just adding a few lines here and there)

---

## How to Use (if you really want to)

```bash
cd src
./make.bash

bin/go build -race2 yourprogram.go
```

The `-race2` flag works exactly like `-race`, except it calls `race2*` functions instead of `race*` functions. The actual race detection algorithm is just stubs - you'd need to implement that yourself.

---

## Why This Approach Sucks

1. **Maintenance nightmare**: Go commits 20-50 changes daily. Keeping this fork in sync is impossible.
2. **Distribution**: Nobody wants to build Go from source just to use your tool.
3. **Overkill**: We duplicated the entire compiler just to change some function names.

**Lesson learned**: Don't fork the compiler. There are better ways to add instrumentation.

---

## The Better Approach

See: **https://github.com/amirkhaki/moriarty**

Instead of forking the compiler, we're trying a different approach that works with standard Go builds.

---

*This was a learning experience. The new repo takes a smarter approach.*
