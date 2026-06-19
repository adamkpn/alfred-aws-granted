#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
make workflow VERSION="${1:-1.0.1}"
