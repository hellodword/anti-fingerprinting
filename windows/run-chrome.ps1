$errorActionPreference='Stop'
$SavePath = "\\host.lan\Data"
$BasePath = $env:TEMP
$Installer = "chrome_installer.exe"
cd $BasePath

# https://stackoverflow.com/a/70736582
switch -File $SavePath\.env{
  default {
    $name, $value = $_.Trim() -split '=', 2
    if ($name -and $name[0] -ne '#') { # ignore blank and comment lines.
      Set-Item "Env:$name" $value
    }
  }
}

# pause windows update
Invoke-WebRequest "https://github.com/vmware/ansible-vsphere-gos-validation/raw/main/windows/utils/scripts/pause_windows_update.ps1" -OutFile C:\pause_windows_update.ps1
powershell.exe -ExecutionPolicy Bypass -File C:\pause_windows_update.ps1

# install cert
certutil -addstore "Root" $SavePath\certs\tls.crt
# add hosts
Add-Content -Path $env:windir\System32\drivers\etc\hosts -Value "`n172.17.0.1`texample.org" -Force
ipconfig /flushdns

# download, verify and install chrome
# Invoke-WebRequest "$env:URL" -OutFile $BasePath\$Installer
# (Get-FileHash $BasePath\$Installer).Hash -eq "$env:HASH"
cp $SavePath\$Installer $BasePath\$Installer
$proc = Start-Process -FilePath $BasePath\$Installer -Args "--silence --install --do-not-launch-chrome --disable-progress" -Verb RunAs -PassThru
$timeouted = $null
$proc | Wait-Process -Timeout 120 -ErrorAction SilentlyContinue -ErrorVariable timeouted

# locate chrome
$chrome = (Get-ChildItem -Path $env:ProgramFiles\Google\Chrome\,$env:LOCALAPPDATA\Google\Chrome\,${env:ProgramFiles(x86)}\Google\Chrome\ -Filter chrome.exe -Recurse -ErrorAction SilentlyContinue -Force).FullName

# $proc = Start-Process -FilePath "$chrome" -Args "-incognito --headless --disable-gpu --print-to-pdf=$SavePath\result.pdf --run-all-compositor-stages-before-draw --ignore-certificate-errors https://example.org:8443/v1/all" -PassThru
# $timeouted = $null
# $proc | Wait-Process -Timeout 120 -ErrorAction SilentlyContinue -ErrorVariable timeouted
# if ($timeouted)
# {
#   $proc | kill
# }

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
