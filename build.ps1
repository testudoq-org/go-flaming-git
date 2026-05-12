# build.ps1 — Cross-compile fgit for Windows, Linux, and macOS.
# Usage: .\build.ps1 [-Version "v1.0.0"]
[CmdletBinding()]
param(
    [string]$Version = "dev"
)

$ErrorActionPreference = "Stop"

$Module   = "github.com/stephenhstewart/fgit/cmd/fgit"
$LdFlags  = "-s -w -X ${Module}.Version=${Version}"
$Out      = "dist"

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
