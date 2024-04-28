$errorActionPreference='Stop'

$SrcPath = "$PWD"

# pause windows update
Invoke-WebRequest "https://github.com/vmware/ansible-vsphere-gos-validation/raw/main/windows/utils/scripts/pause_windows_update.ps1" -OutFile C:\pause_windows_update.ps1
powershell.exe -ExecutionPolicy Bypass -File C:\pause_windows_update.ps1

# install cert
certutil -addstore "Root" $SrcPath\certs\tls.crt
# add hosts
Add-Content -Path $env:windir\System32\drivers\etc\hosts -Value "`n172.17.0.1`texample.org" -Force
ipconfig /flushdns
