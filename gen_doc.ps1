Remove-Item -Path .\docs\* -Recurse -Force

Set-Location .\client

cargo clean

cargo doc --no-deps

'<meta http-equiv="refresh" content="0; url=client/index.html">' | Set-Content -Path .\target\doc\index.html

Copy-Item -Path .\target\doc\* -Destination ..\docs -Recurse -Force

cargo clean

Set-Location ..
