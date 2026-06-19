#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

ROOT="assets/source"
mkdir -p assets/source icons

if [[ ! -f assets/source/icon.png ]]; then
  cp -f icon.png assets/source/icon.png 2>/dev/null || true
fi
if [[ ! -f assets/source/region.png ]]; then
  cp -f region.png assets/source/region.png 2>/dev/null || true
fi

round_image() {
  local src="$1"
  local dst="$2"
  local size="$3"
  local radius="$4"

  magick "$src" -resize "${size}x${size}" \
    \( +clone -alpha extract -draw "roundrectangle 0,0 ${size},${size} ${radius},${radius}" \) \
    -compose CopyOpacity -composite "$dst"
}

round_circle() {
  local src="$1"
  local dst="$2"
  local size="$3"

  magick "$src" -resize "${size}x${size}" \
    \( +clone -alpha extract -size "${size}x${size}" xc:none \
      -fill white -draw "circle $((size / 2)),$((size / 2)) $((size / 2)),0" \) \
    -alpha off -compose CopyOpacity -composite "$dst"
}

# Workflow + profiles: fully round AWS logo.
round_circle "${ROOT}/icon.png" "icon.png" 256
round_circle "${ROOT}/icon.png" "icons/profile.png" 256

# Regions: strongly rounded squircle.
round_image "${ROOT}/region.png" "region.png" 256 72

echo "Generated rounded icons: icon.png, icons/profile.png, region.png"
