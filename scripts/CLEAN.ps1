Set-Location ..

if (Test-Path -Path release) {
    Remove-Item -Recurse -Force release
}

$temp="C:\$Env:USERNAME\AppData\Local\go-build"
if (Test-Path -Path $temp) {
    Remove-Item -Recurse -Force $temp
}

Set-Location scripts
exit 0
