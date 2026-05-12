# build.ps1 — Cross-compile fgit for Windows, Linux, and macOS.
# Usage: .\build.ps1 [-Version "v1.0.0"]
[CmdletBinding()]
param(
    [string]$Version = "dev",
    [switch]$CrapGate,
    [int]$CrapThreshold = 15
)

$ErrorActionPreference = "Stop"

$Module   = "github.com/stephenhstewart/fgit/cmd/fgit"
$LdFlags  = "-s -w -X ${Module}.Version=${Version}"
$Out      = "dist"

if ($CrapGate) {
    Write-Host "Running quality gate with CRAP threshold $CrapThreshold"

    go test ./... -coverprofile=coverage.out
    go run ./cmd/cover2lcov -in coverage.out -out coverage.lcov
    go vet ./...
    go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

    $tmpRepo = Join-Path $env:TEMP ('go-crap4go-' + [guid]::NewGuid().ToString('N'))
    git clone --depth 1 https://github.com/testudoq-org/go-crap4go.git $tmpRepo | Out-Null
    Push-Location $tmpRepo
    go build -o crap.exe ./cmd/crap
    Pop-Location

    & (Join-Path $tmpRepo 'crap.exe') cmd/fgit internal/alias internal/dispatch internal/runner --no-run-tests --coverprofile=coverage.out --threshold=$CrapThreshold | Tee-Object -FilePath crap-report.txt

    if ($LASTEXITCODE -ne 0) {
        Write-Error "CRAP gate failed. Investigate functions at or above threshold in crap-report.txt"
        exit 1
    }

    Write-Host "Quality gate passed"
    exit 0
}

New-Item -ItemType Directory -Force -Path $Out | Out-Null

$Targets = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; Binary = "fgit-windows-amd64.exe" },
    @{ GOOS = "linux";   GOARCH = "amd64"; Binary = "fgit-linux-amd64"       },
    @{ GOOS = "linux";   GOARCH = "arm64"; Binary = "fgit-linux-arm64"       },
    @{ GOOS = "darwin";  GOARCH = "amd64"; Binary = "fgit-darwin-amd64"      },
    @{ GOOS = "darwin";  GOARCH = "arm64"; Binary = "fgit-darwin-arm64"      }
)

foreach ($t in $Targets) {
    $output = Join-Path $Out $t.Binary
    Write-Host "Building $($t.GOOS)/$($t.GOARCH) → $output"
    $env:GOOS   = $t.GOOS
    $env:GOARCH = $t.GOARCH
    go build -trimpath -ldflags $LdFlags -o $output .
    Remove-Item Env:\GOOS
    Remove-Item Env:\GOARCH
}

Write-Host ""
Write-Host "Build complete. Binaries in .\$Out\"
Get-ChildItem $Out | Select-Object Name, Length
