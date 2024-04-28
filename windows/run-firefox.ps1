$errorActionPreference='Stop'
$SavePath = "$PWD"
$BasePath = $env:TEMP
$Installer = "firefox_installer.exe"
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

cp $SavePath\$Installer $BasePath\$Installer
$proc = Start-Process -FilePath $BasePath\$Installer -Args "/S" -Verb RunAs -PassThru
$timeouted = $null
$proc | Wait-Process -Timeout 120 -ErrorAction SilentlyContinue -ErrorVariable timeouted

# locate firefox
$firefox = (Get-ChildItem -Path "$env:ProgramFiles\Mozilla Firefox\" -Filter firefox.exe -Recurse -ErrorAction SilentlyContinue -Force).FullName

# # 00000000-0000-0000-0000-000000000001 incognito / private
# # 00000000-0000-0000-0000-000000000002 normal
# # 00000000-0000-0000-0000-000000000003 normal disable http2

Start-Sleep -Seconds 3
Start-Process -FilePath "$firefox" -Args "-private https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000001" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
Start-Process -FilePath "$firefox" -Args "-profile $env:TEMP\firefox2-$_uuid -CreateProfile ""firefox2-$_uuid $env:TEMP\firefox2-$_uuid"" https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000002" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
mkdir $env:TEMP\firefox3-$_uuid
echo 'user_pref("network.http.http2.enabled", false);' | Out-File -encoding ASCII $env:TEMP\firefox3-$_uuid\user.js
Start-Process -FilePath "$firefox" -Args "-profile $env:TEMP\firefox3-$_uuid -CreateProfile ""firefox3-$_uuid $env:TEMP\firefox3-$_uuid"" https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000003" -PassThru
