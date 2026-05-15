# Contributing to contctrl

Thanks for your interest in contributing. This is a focused tool with a simple scope — contributions that keep it minimal and practical are most welcome.

## What's welcome

- Bug fixes
- New CLI commands (`restart`, `logs`, etc.)
- Config file support for `BASE_DIR` and other settings
- Improvements to fuzzy selection flow
- Better error handling and messaging
- Documentation improvements

## What to avoid

- Heavy dependencies or TUI frameworks
- Features that change the core fuzzy-first navigation model
- Breaking changes to the `metadata.yml` format without discussion

## Getting started

```sh
git clone https://github.com/hanymamdouh82/contctrl.git
cd contctrl
go mod tidy
go build ./...
```

Make sure you have Docker with the Compose plugin installed to test end-to-end.

## Project layout

```
.
├── main.go                  # Entry point and command dispatch
├── internal/
│   ├── navigator/           # Project catalog builder
│   ├── selector/            # Fuzzy finder wrappers
│   └── ui/                  # Terminal output helpers
```

## Workflow

1. Fork the repo and create a branch from `main`
2. Make your changes
3. Test manually against a real Compose project directory
4. Open a pull request with a clear description of what and why

## Pull request guidelines

- Keep PRs focused — one concern per PR
- Include a short description of the problem you're solving
- If adding a new command, follow the existing pattern in `main.go`
- No new dependencies without prior discussion in an issue

## Reporting issues

Open a GitHub issue with:

- What you ran
- What you expected
- What actually happened
- Your OS, Go version, and Docker version

## Questions

Open a GitHub issue with the `question` label.
