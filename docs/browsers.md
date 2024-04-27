## Browsers

> This file is only for recording some useful resources

### Desktop

- Chrome

  - ~~https://googlechromelabs.github.io/chrome-for-testing/known-good-versions-with-downloads.json~~ This is chromium
  - ~~https://github.com/browser-actions/setup-chrome~~
  - https://stackoverflow.com/questions/54927496/how-to-download-older-versions-of-chrome-from-a-google-official-site
    ```sh
    curl -d '<?xml version="1.0" encoding="UTF-8"?>
    <request protocol="3.0" updater="Omaha" sessionid="$(uuidgen)"
        installsource="update3web-ondemand" requestid="$(uuidgen)">
        <os platform="win" version="10.0" arch="x64" />
        <app appid="{8A69D345-D564-463C-AFF1-A69D9E530F96}" version="5.0.375"
            ap="x64-stable-statsdef_0" lang="" brand="GCEB">
            <updatecheck targetversionprefix="124"/>
        </app>
    </request>' -X POST https://tools.google.com/service/update2 \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -H 'X-Goog-Update-Interactivity: fg'
    ```
  - https://github.com/ScoopInstaller/Extras/commits/master/bucket/googlechrome.json
    - https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/googlechrome.json&sha=master&per_page=100&page=1
    ```sh
    curl -sSL https://github.com/ScoopInstaller/Extras/raw/3261258440d54ac95955687d8a2d40c615337d5c/bucket/googlechrome.json | \
        jq -r '.architecture."64bit".url' | \
        awk -F# '{print $1}'
    ```
  - https://github.com/chocolatey-community/chocolatey-packages/commits/master/automatic/googlechrome/googlechrome.nuspec
  - https://github.com/microsoft/winget-pkgs/commits/master/manifests/g/Google/Chrome
  - https://web.archive.org/web/20220623192848/https://dl.google.com/release2/chrome/adpllftgiudg2qsog7pdou2vggiq_103.0.5060.53/103.0.5060.53_chrome_installer.exe
  - https://github.com/chromium/chromium/blob/821f200e1299db31eefaabbf731b2e1b4d7c71da/chrome/installer/util/util_constants.cc

    ```
    installer.exe --verbose-logging --do-not-launch-chrome --channel=stable --allow-downgrade --trigger-active-setup
    ```

    ```
    chrome.exe --user-data-dir=/tmp/xxx --ignore-certificate-errors
    chrome.exe -incognito --ignore-certificate-errors
    chrome.exe --user-data-dir=/tmp/xxx --disable-http2 --ignore-certificate-errors
    chrome.exe --disable-http2 -incognito --ignore-certificate-errors

    C:\Users\admin\AppData\Local\Google\Chrome\Application\chrome.exe -incognito --headless --disable-gpu --print-to-pdf=$env:TEMP\test.pdf --run-all-compositor-stages-before-draw --ignore-certificate-errors https://192.168.0.2:8443/v1/all
    ```

- Safari

  - Docker-OSX
  - https://github.com/search?q=%2Fsafaridriver%2F+%28path%3A.github%2Fworkflows%2F*.yml+OR+path%3A.github%2Fworkflows%2F*.yaml%29&type=code
  - https://developer.apple.com/documentation/webkit/testing_with_webdriver_in_safari
  - https://www.selenium.dev/documentation/webdriver/browsers/safari/
  - https://appium.io/docs/zh/latest/
  - WebKit: https://playwright.dev/docs/browsers

- Edge

  - https://github.com/browser-actions/setup-edge
  - https://www.microsoft.com/en-us/edge/business/download
  - https://www.catalog.update.microsoft.com/Search.aspx?q=Microsoft%20Edge%20

  ```sh
  wget https://catalog.s.download.windowsupdate.com/d/msdownload/update/software/updt/2024/04/microsoftedgeenterprisex64_27649128f9d43e8e965ead806b80e367d3582c64.cab
  cabextract microsoftedgeenterprisex64_27649128f9d43e8e965ead806b80e367d3582c64.cab
  file MicrosoftEdgeEnterpriseX64.msi
  ```

  - https://github.com/he3als/EdgeRemover
  - https://learn.microsoft.com/en-us/answers/questions/1359115/how-can-i-roll-back-a-edge-version-from-115-to-113
  - https://learn.microsoft.com/en-us/deployedge/edge-learnmore-rollback
  - https://pureinfotech.com/rollback-previous-version-microsoft-edge/
  - https://superuser.com/questions/1759206/how-to-install-rollback-to-a-specific-version-of-microsoft-edge-on-windows
  - https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/msiexec#parameters-1

- Firefox

  - https://github.com/browser-actions/setup-firefox
  - https://ftp.mozilla.org/pub/firefox/releases/

  ```sh
  curl -fsSL https://ftp.mozilla.org/pub/firefox/releases/ | grep -oP '(?<="/pub/firefox/releases/)\d+(\.\d+)+(?=/")' | sort -Vr | head -n 10
  ```

  - https://edgeupdates.microsoft.com/api/products
  - https://github.com/he3als/EdgeRemover/blob/5a7b3b9fc1891f071ce1bde43c6715a64c6a928a/RemoveEdge.ps1#L254
  - https://firefox-source-docs.mozilla.org/browser/installer/windows/installer/FullConfig.html
  - https://wiki.mozilla.org/Firefox/CommandLineOptions
  - https://stackoverflow.com/a/77009337

### Mobile

- Chrome & Samsung Internet
  - Android emulator
    - https://developer.android.com/studio/test/gradle-managed-devices
      > `Note: When using Gradle-managed devices on servers that don't support hardware rendering, such as GitHub Actions, you need to specify the following flag: -Pandroid.testoptions.manageddevices.emulator.gpu=swiftshader_indirect. If you use PowerShell, you have to surround the flag with quotes.`
      - https://github.com/search?q=%2Fandroid.testoptions.manageddevices.emulator.gpu%2F+%28path%3A.github%2Fworkflows%2F*.yml+OR+path%3A.github%2Fworkflows%2F*.yaml%29+NOT+path%3A.github%2Fworkflows%2FAndroidCIWithGmd.yaml&type=code
    - https://github.com/actions/runner-images/issues/8676
    - https://github.com/orgs/community/discussions/70344
    - https://github.com/android/nowinandroid/blob/ead3f49f7bf94f9ef0be1256098554bac51f4a7c/.github/workflows/Release.yml
    - https://gitlab.com/newbit/rootAVD
    - https://automationchronicles.com/error-when-opening-chrome-on-android-13-via-adb/
  - Real device testing: Firebase/testgrid/testdevlab/browserstack
- Safari
  - Docker-OSX + emulator + appium
  - corellium + appium
