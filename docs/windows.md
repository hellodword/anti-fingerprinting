## windows

I choose to implement it with `dockur/windows`(https://github.com/dockur/windows/issues/187#issuecomment-1960638868), not the GitHub Actions, because I want to make sure it'll be working anywhere, and I believe the full featured windows will provide more useful information.

### cmd/ps1 snippets

- turn on Network Discovery and File Sharing

```ps1
netsh advfirewall firewall set rule group="Network Discovery" new enable=Yes
```

- OpenSSH
  - https://github.com/containerd/containerd/blob/9d108fa83bfa8c477e233d6729f3b0dba96cfff6/script/setup/enable_ssh_windows.ps1
  - https://github.com/microsoft/AirSim/blob/6688d27d3712c2a9c824ababec7a2703475b6628/azure/azure-env-creation/configure-vm.ps1#L9-L15
  - https://github.com/containerd/containerd/blob/9d108fa83bfa8c477e233d6729f3b0dba96cfff6/integration/images/README.md#configure-needed-services
  - https://github.com/search?q=%2FOpenSSH.Server%7E%7E%7E%7E0.0.1.0%2F&type=code

```ps1
# # Open Firewall port for ssh
# # New-NetFirewallRule : Cannot create a file when that file already exists.
# New-NetFirewallRule -Name 'OpenSSH-Server-In-TCP' -DisplayName 'OpenSSH Server (sshd)' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
```

- Install Chocolatey

  - https://github.com/microsoft/AirSim/blob/6688d27d3712c2a9c824ababec7a2703475b6628/azure/azure-env-creation/configure-vm.ps1#L17C2-L25

- param with space must be quoted by `'`

```sh
# directly execute
powershell -NoLogo -NonInteractive -ExecutionPolicy ByPass -File C:/Users/Docker/Downloads/enable_ssh_windows.ps1 -SSHPublicKey '123 456'

# via ssh
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i /path/to/sshkey docker@127.0.0.1 -p 2222 \
  powershell -NoLogo -NonInteractive -ExecutionPolicy ByPass -File C:/Users/Docker/Downloads/enable_ssh_windows.ps1 -SSHPublicKey "'123 456'"
```

or read from file

```xml
<SynchronousCommand>
  <Order>17</Order>
  <CommandLine>powershell.exe -sta -ExecutionPolicy Unrestricted -Command "C:\Users\Docker\Downloads\enable_ssh_windows.ps1 -SSHPublicKey (cat \\host.lan\Data\dockur.pub) *> C:\output-17.log"</CommandLine>
  <Description>Enable ssh</Description>
</SynchronousCommand>
```

- Download file

```xml
<SynchronousCommand>
  <Order>16</Order>
  <CommandLine>powershell.exe -NoLogo -Command "(new-object System.Net.WebClient).DownloadFile('https://raw.githubusercontent.com/containerd/containerd/9d108fa83bfa8c477e233d6729f3b0dba96cfff6/script/setup/enable_ssh_windows.ps1', 'C:\Users\Docker\Downloads\enable_ssh_windows.ps1')"</CommandLine>
  <Description>Download ssh ps1</Description>
</SynchronousCommand>
```

- find file from multiple paths

```ps1
(Get-ChildItem -Path $env:ProgramFiles\Google\Chrome\Application\chrome.exe,$env:LOCALAPPDATA\Google\Chrome\Application\chrome.exe -Filter chrome.exe -Recurse -ErrorAction SilentlyContinue -Force).FullName
```

- Write file

`echo $content > $filepath` will be `UTF-16 LE`, use `echo 'user_pref("network.http.http2.enabled", false);' | Out-File -encoding ASCII $env:TEMP\firefox3\user.js` instead
