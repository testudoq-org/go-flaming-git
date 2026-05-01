#!/usr/bin/env bash
# build.sh — cross-compile fgit for Windows, Linux, and macOS.
# Usage: ./build.sh [version]
#   version defaults to "dev" when not supplied.

set -euo pipefail

VERSION="${1:-dev}"
LDFLAGS="-s -w -X github.com/stephenhstewart/fgit/cmd/fgit.Version=${VERSION}"
OUT="dist"

mkdir -p "${OUT}"

targets=(
  "windows/amd64/fgit.exe"
  "linux/amd64/fgit"
  "linux/arm64/fgit"
  "darwin/amd64/fgit"
  "darwin/arm64/fgit"
)

for target in "${targets[@]}"; do
  IFS='/' read -r GOOS GOARCH BINARY <<< "${target}"
  output="${OUT}/${BINARY}-${GOOS}-${GOARCH}"
  if [[ "${GOOS}" == "windows" ]]; then
    output="${OUT}/fgit-${GOOS}-${GOARCH}.exe"
  fi
  echo "Building ${GOOS}/${GOARCH} → ${output}"
  GOOS="${GOOS}" GOARCH="${GOARCH}" go build \
    -trimpath \
    -ldflags="${LDFLAGS}" \
    -o "${output}" \
    .
done

echo ""
echo "Build complete. Binaries in ./${OUT}/"
ls -lh "${OUT}/"
