$Env:GOROOT_FINAL="/usr"

$Env:CGO_ENABLED=0
$Env:GOOS="linux"
$Env:GOARCH="amd64"
go build -a -trimpath -asmflags "-s -w" -ldflags "-s -w" -gcflags "-l=4" -o release\stream
if (! $?) { exit 1 }

Copy-Item -Force "default.json" release
Copy-Item -Force scripts\installer\*.sh release

exit 0
