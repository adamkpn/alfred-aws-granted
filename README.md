# Alfred AWS Granted (Go)

[![CI](https://github.com/adamkpn/alfred-aws-granted/actions/workflows/ci.yml/badge.svg)](https://github.com/adamkpn/alfred-aws-granted/actions/workflows/ci.yml)

Alfred workflow to open AWS Console via [Granted](https://github.com/fwdcloudsec/granted), with **fuzzy search** for profiles, services, and regions.

Inspired by [sherifabdlnaby/alfred-aws-granted](https://github.com/sherifabdlnaby/alfred-aws-granted), rewritten in Go using [awgo](https://github.com/deanishe/awgo) for Alfred integration and fuzzy matching.

## Features

- Fuzzy-search AWS profiles from `~/.aws/config` and `~/.aws/credentials`
- All AWS console services supported by Granted's `-s` flag
- Fuzzy-search AWS regions
- Keyboard shortcuts:
  - **Return** — select profile → service → open console
  - **⌘ Return** — open console home (skip service selection)
  - **⌥ Return** — select region first

## Prerequisites

1. [Granted](https://granted.dev) installed and configured:

```bash
brew tap fwdcloudsec/granted
brew install granted
```

2. Follow Granted's [Getting Started](https://docs.granted.dev/getting-started/) guide.

## Installation

### From GitHub Releases (recommended)

1. Download the latest `alfred-granted-v*.alfredworkflow` from [Releases](https://github.com/adamkpn/alfred-aws-granted/releases).
2. Double-click the file to import it into Alfred.
3. **macOS Gatekeeper:** the workflow includes a compiled binary. After importing, remove the download quarantine once:
   - Alfred → Preferences → Workflows → **AWS Granted** → **⋮** → **Open in Finder**
   - In Terminal, from that folder:

```bash
./scripts/trust-workflow.sh
```

   Or manually: `xattr -cr .` in the workflow folder opened from Alfred.

### From source

Requires Go 1.22+, ImageMagick (`brew install imagemagick`), and macOS (for universal binary build):

```bash
make workflow
open build/alfred-granted-v1.0.1.alfredworkflow
```

### Development

```bash
make build    # universal macOS binary (local only, not committed)
make test
./alfred-aws-granted profiles "prod"
./alfred-aws-granted services "lambda"
./alfred-aws-granted regions "us-east"
```

Link the workflow directory in Alfred Preferences for live development.

## Usage

1. Open Alfred and type `aws` (configurable keyword).
2. Fuzzy-search your AWS profile and press Return.
3. Pick a service, or use **⌘** to open the console home directly.
4. Use **⌥** on the profile step to choose a region.

## macOS Gatekeeper

If Alfred shows *"Apple could not verify aws-granted-profiles is free of malware"*, macOS is blocking the unsigned workflow binary downloaded from GitHub. This is expected for open-source Alfred workflows without Apple notarization.

**Fix (one time per install):**

1. Alfred → Preferences → Workflows → **AWS Granted** → **⋮** → **Open in Finder**
2. Run `./scripts/trust-workflow.sh` in Terminal from that folder

Alternatively, build and install from source — locally built binaries are not quarantined.

## Releasing

Push a semver tag to trigger the release workflow, which builds the `.alfredworkflow` on macOS and attaches it to a GitHub Release:

```bash
git tag v1.0.1
git push origin v1.0.1
```

Binaries and build artifacts are **not** committed — they are produced in CI and published as release assets only.

## Project layout

```
cmd/alfred-aws-granted/     # Go binary entrypoint (profiles, services, regions, open)
internal/profiles/          # AWS profile discovery
internal/services/          # Granted console services
internal/regions/           # AWS regions
internal/granted/           # assume -c wrapper
services/                   # AWS service icons (SVG)
info.plist                  # Alfred workflow definition
```

## License

MIT

## Credits

- [Granted](https://granted.dev) by fwd:cloudsec
- [alfred-aws-granted](https://github.com/sherifabdlnaby/alfred-aws-granted) by Sherif Abdel-Naby
- [awgo](https://github.com/deanishe/awgo) by Dean Jackson
