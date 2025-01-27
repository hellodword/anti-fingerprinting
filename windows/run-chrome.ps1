$errorActionPreference='Stop'

$SrcPath = "$PWD"
$Installer = "chrome_installer.exe"

# https://stackoverflow.com/a/70736582
switch -File $SrcPath\.env{
  default {
    $name, $value = $_.Trim() -split '=', 2
    if ($name -and $name[0] -ne '#') { # ignore blank and comment lines.
      Set-Item "Env:$name" $value
    }
  }
}

$proc = Start-Process -FilePath $SrcPath\$Installer -Args "--silence --install --do-not-launch-chrome --disable-progress" -Verb RunAs -PassThru
$timeouted = $null
$proc | Wait-Process -Timeout 120 -ErrorAction SilentlyContinue -ErrorVariable timeouted

# locate chrome
$chrome = (Get-ChildItem -Path $env:ProgramFiles\Google\Chrome\,$env:LOCALAPPDATA\Google\Chrome\,${env:ProgramFiles(x86)}\Google\Chrome\ -Filter chrome.exe -Recurse -ErrorAction SilentlyContinue -Force).FullName

# 00000000-0000-0000-0000-000000000001 incognito
# 00000000-0000-0000-0000-000000000002 normal
# 00000000-0000-0000-0000-000000000003 normal disable http2

Start-Sleep -Seconds 3
Start-Process -FilePath "$chrome" -Args "-incognito https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000001" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
Start-Process -FilePath "$chrome" -Args "--no-default-browser-check --no-first-run --user-data-dir=$env:TEMP\chrome2-$_uuid https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000002" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
Start-Process -FilePath "$chrome" -Args "--no-default-browser-check --no-first-run --user-data-dir=$env:TEMP\chrome3-$_uuid https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000003 --disable-http2" -PassThru
