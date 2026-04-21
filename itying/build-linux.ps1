$ErrorActionPreference = "Stop"

Write-Host "Building Linux binary..." -ForegroundColor Green

$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"

go build -ldflags="-s -w" -o admin_linux ./main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build success! Output: admin_linux" -ForegroundColor Cyan
} else {
    Write-Host "Build failed. Check error above." -ForegroundColor Red
}