#!/usr/bin/env bash
set -euo pipefail

# Remove Gatekeeper quarantine from the installed workflow directory.
# Alfred blocks unsigned downloaded binaries until quarantine is cleared.
#
# Usage (after importing the workflow):
#   1. Alfred → Preferences → Workflows → AWS Granted → ⋮ → Open in Finder
#   2. In Terminal: ./scripts/trust-workflow.sh

root="$(cd "$(dirname "$0")/.." && pwd)"

echo "Removing Gatekeeper quarantine from:"
echo "  $root"
xattr -cr "$root"
echo "Done. Retry the workflow in Alfred."
