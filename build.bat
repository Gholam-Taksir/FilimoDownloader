@echo off
title Building FilimoDownloader
color 0E
echo ==========================================
echo    Building FilimoDownloader v2.0.0
echo ==========================================
echo.

go build -ldflags="-X main.isProduction=true" -o FilimoDownloader.exe ./cmd/FilimoDownloader

if %errorlevel% equ 0 (
    echo.
    echo ==========================================
    echo    Build successful!
    echo    FilimoDownloader.exe is ready
    echo ==========================================
) else (
    echo.
    echo ==========================================
    echo    Build failed! Check errors above
    echo ==========================================
)
pause
