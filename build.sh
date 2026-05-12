#!/usr/bin/env bash
# build.sh — cross-compile fgit for Windows, Linux, and macOS.
# Usage: ./build.sh [version]
#   version defaults to "dev" when not supplied.

set -euo pipefail

if [[ "${1:-}" == "crap" ]]; then
  THRESHOLD="${2:-15}"

  echo "Running quality gate with CRAP threshold ${THRESHOLD}"
  go test ./... -coverprofile=coverage.out
  go run ./cmd/cover2lcov -in coverage.out -out coverage.lcov
  go vet ./...
  go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

  tmp_repo="$(mktemp -d)"
  git clone --depth 1 https://github.com/testudoq-org/go-crap4go.git "${tmp_repo}" >/dev/null
  (
    cd "${tmp_repo}"
    go build -o crap ./cmd/crap
  )

  "${tmp_repo}/crap" cmd/fgit internal/alias internal/dispatch internal/runner --no-run-tests --coverprofile=coverage.out --threshold="${THRESHOLD}" | tee crap-report.txt

  echo "Quality gate passed"
  exit 0
fi

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
