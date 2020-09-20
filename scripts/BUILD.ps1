Set-Location ..

if (Test-Path -Path release) {
    Remove-Item -Recurse -Force release
}

$temp="C:\$Env:USERNAME\AppData\Local\go-build"
if (Test-Path -Path $temp) {
    Remove-Item -Recurse -Force $temp
}

$Env:CGO_ENABLED="0"
$Env:GOROOT_FINAL="/usr"

$Env:GOOS="linux"
$Env:GOARCH="amd64"
go build -a -trimpath -asmflags "-s -w" -ldflags "-s -w" -gcflags "-l=4" -o release\stream
if (! $?) { exit 1 }

upx --ultra-brute release\stream

Copy-Item -Force "default.json" release
Copy-Item -Force scripts\installer\*.sh release

Set-Location release
Compress-Archive -Path * -DestinationPath release.zip

certutil -hashfile stream SHA256
certutil -hashfile release.zip SHA256

Set-Location ..\scripts
exit 0
