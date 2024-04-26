$errorActionPreference='Stop'
$SavePath = "\\host.lan\Data"
$BasePath = $env:TEMP
$Installer = "edge_installer.exe"
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

# disable edge auto update
New-NetFirewallRule -DisplayName "BlockEdgeUpdate" -Direction Outbound -Action Block -Program "${env:ProgramFiles(x86)}\Microsoft\EdgeUpdate\MicrosoftEdgeUpdate.exe"

$edge="${env:ProgramFiles(x86)}\Microsoft\Edge\Application\msedge.exe"

$oriProductVersion=(Get-Item "$edge").VersionInfo.ProductVersion

cp $SavePath\$Installer $BasePath\$Installer.msi
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$BasePath\$Installer.msi`" /quiet ALLOWDOWNGRADE=1" -Wait


(Get-Item $edge).VersionInfo

# trigger unknown setup by starting edge after upgrading/downgrading

$timeout = 600
$startTime = Get-Date
$success = $null
while ((Get-Date) -lt ($startTime.AddSeconds($timeout))) {

  Start-Process -FilePath "$edge" -PassThru
  Start-Sleep -Seconds 10
  Stop-Process -Name msedge -Force
  Start-Sleep -Seconds 3

  $productVersion = (Get-Item "$edge").VersionInfo.ProductVersion
  $fileVersion = (Get-Item "$edge").VersionInfo.FileVersion
  Write-Host "productVersion $productVersion fileVersion $fileVersion"
  if ($productVersion -eq "$env:BROWSER_VERSION" -and $fileVersion -eq "$env:BROWSER_VERSION") {
    $success = "true"
    break
  }
}

if (!$success) {
  Throw "replace $oriProductVersion with $env:BROWSER_VERSION failed"
}

# 00000000-0000-0000-0000-000000000001 incognito / inprivate
# 00000000-0000-0000-0000-000000000002 normal
# 00000000-0000-0000-0000-000000000003 normal disable http2
Start-Sleep -Seconds 3
Start-Process -FilePath "$edge" -Args "-inprivate https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000001" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
Start-Process -FilePath "$edge" -Args "--no-default-browser-check --no-first-run --user-data-dir=$env:TEMP\edge2-$_uuid https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000002" -PassThru

Start-Sleep -Seconds 3
$_uuid=([guid]::NewGuid().ToString())
Start-Process -FilePath "$edge" -Args "--no-default-browser-check --no-first-run --user-data-dir=$env:TEMP\edge3-$_uuid https://example.org:8443/v1/all?id=00000000-0000-0000-0000-000000000003 --disable-http2" -PassThru
