setlocal
set /P "VERSION=version ? "
for /F %%I in ('cd') do set NAME=%%~nI
for %%I in (386 amd64) do call :build %%I
endlocal
exit /b 0

:build
    set GOARCH=%1
    go build -ldflags "-s -w"
    upx -9 %NAME%.exe
    zip -9 %NAME%-windows-%1-%VERSION%.zip %NAME%.exe
    exit /b
