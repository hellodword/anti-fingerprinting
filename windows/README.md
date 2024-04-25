## windows

### windows in docker

- https://github.com/dockur/windows/issues/187
- Authentication: https://github.com/dockur/windows/issues/301#issuecomment-2018610554

```sh
# .htpasswd
# openssl passwd -apr1

curl -fsSL --output ./windows/nginx.conf https://github.com/qemus/qemu-docker/raw/master/web/nginx.conf
sed -i '2s@^@    auth_basic "Administrator’s Area";\n    auth_basic_user_file /storage/shared/.htpasswd;\n@' ./windows/nginx.conf

# https://unix.stackexchange.com/a/69318
ssh-keygen -t rsa -b 4096 -C "dockur" -f dockur-sshkey -q -N ""
mv ./dockur-sshkey.pub ./windows/shared/
sudo chmod 600 ./dockur-sshkey

mkdir -p ./windows/shared/certs
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -days 3650 \
  -nodes -keyout ./windows/shared/certs/tls.key -out ./windows/shared/certs/tls.crt -subj "/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,DNS:*.localhost,DNS:example.org,IP:127.0.0.1,IP:172.17.0.1"

cp ./windows/run-chrome.ps1 ./windows/shared/run-custom.ps1

docker run --rm -it \
  --stop-timeout 120 --name windows \
  -e MANUAL=N \
  -v /path/to/Win11_23H2_English_x64v2.iso:/storage/custom.iso:ro \
  -v $PWD/windows/shared:/storage/shared:rw \
  -v $PWD/windows/win11x64.xml:/run/assets/win11x64.xml:ro \
  -v $PWD/windows/nginx.conf:/etc/nginx/sites-enabled/web.conf:ro \
  -p 127.0.0.1:2222:22 -p 127.0.0.1:3389:3389 -p 127.0.0.1:8006:8006 \
  --device=/dev/kvm --cap-add NET_ADMIN dockurr/windows

docker run --rm -it \
  --stop-timeout 120 --name windows \
  -e MANUAL=N \
  -e VERSION=win11 \
  -v $PWD/windows/shared:/storage/shared:rw \
  -v $PWD/windows/win11x64.xml:/run/assets/win11x64.xml:ro \
  -v $PWD/windows/nginx.conf:/etc/nginx/sites-enabled/web.conf:ro \
  -p 127.0.0.1:2222:22 -p 127.0.0.1:3389:3389 -p 127.0.0.1:8006:8006 \
  --device=/dev/kvm --cap-add NET_ADMIN dockurr/windows

docker run --rm -it \
  --stop-timeout 120 --name windows \
  -e MANUAL=N \
  -v /path/to/Win10_22H2_English_x64v1.iso:/storage/custom.iso:ro \
  -v $PWD/windows/shared:/storage/shared:rw \
  -v $PWD/windows/win10x64.xml:/run/assets/win10x64.xml:ro \
  -v $PWD/windows/nginx.conf:/etc/nginx/sites-enabled/web.conf:ro \
  -p 127.0.0.1:2222:22 -p 127.0.0.1:3389:3389 -p 127.0.0.1:8006:8006 \
  --device=/dev/kvm --cap-add NET_ADMIN dockurr/windows

ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i $PWD/dockur-sshkey docker@127.0.0.1 -p 2222 ls
```

### ps1

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
