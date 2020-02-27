setlocal
set GOARCH=386
go build -ldflags "-s -w" && upx -9 *.exe
endlocal
exit /b


