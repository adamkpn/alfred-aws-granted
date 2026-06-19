#!/bin/bash
set -euo pipefail

export PATH="/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:${PATH:-}"

if [[ -z "${HOME:-}" || ! -f "${HOME}/.aws/config" ]]; then
  user="$(/usr/bin/id -un 2>/dev/null || true)"
  if [[ -n "$user" ]]; then
    detected_home="$(/usr/bin/dscl . -read "/Users/${user}" NFSHomeDirectory 2>/dev/null | /usr/bin/awk '{print $2}')"
    if [[ -n "$detected_home" ]]; then
      export HOME="$detected_home"
    fi
  fi
fi

exec "$(dirname "$0")/../aws-granted-services" "${1:-}"
