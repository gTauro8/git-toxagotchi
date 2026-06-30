# Security Policy

## Read-Only Guarantee

Git-Toxagotchi's core is **strictly read-only**:

- Never modifies Git history, commits, remotes, or `.mailmap`
- Never writes outside `~/.config/git-toxagotchi/` and `~/.local/share/git-toxagotchi/`
- Never makes network requests in the core MVP
- All "chaos" is simulated text animations only

## Pre-Commit Hook

The optional pre-commit hook (`git-toxagotchi hook install`):

- Runs `git diff --cached` (read-only)
- Blocks commits **only** if secrets are detected in staged files
- All other issues are warnings only
- Hook can be uninstalled at any time: `git-toxagotchi hook uninstall`
- Hook script is installed to `.git/hooks/pre-commit` — inspectable plain text

## Secret Detection

The analyzer scans staged diffs for common secret patterns:

- AWS Access Key IDs (`AKIA...`)
- Private key headers (`-----BEGIN ... PRIVATE KEY-----`)
- Generic tokens (`token`, `password`, `secret`, `api_key` in assignment context)

**False positives**: The detection is heuristic. It may flag test data or example values. Override with `--no-verify` when you're certain it's safe (this bypasses the hook entirely, which is standard Git behavior).

## Reporting Vulnerabilities

Please report security issues to: giuseppe.tauro00@gmail.com

Do not open public GitHub issues for security vulnerabilities.

## Supported Versions

Only the latest release is supported.
