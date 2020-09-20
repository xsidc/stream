Set-Location release
Compress-Archive -Path * -DestinationPath release.zip
Move-Item -Force release.zip ..
Set-Location ..
