# TLS fingerprinting

## Browsers

### Desktop

- https://github.com/dockur/windows/issues/187
- https://github.com/dockur/windows/issues/212

- Chrome

  - https://googlechromelabs.github.io/chrome-for-testing/
  - https://github.com/browser-actions/setup-chrome
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

- Firefox
  - https://github.com/browser-actions/setup-firefox
  - https://ftp.mozilla.org/pub/firefox/releases/

### Mobile

- Chrome & Samsung Internet
  - Android emulator
  - Real device testing: Firebase/testgrid/testdevlab/browserstack
- Safari
  - Docker-OSX + emulator + appium
  - corellium + appium

---

## Links

- https://github.com/salesforce/ja3
- https://github.com/FoxIO-LLC/ja4
- https://github.com/refraction-networking/uquic
- https://github.com/gaukas/clienthellod
- https://github.com/dreadl0ck/ja3

- [JA 指纹识全系讲解（上）](https://github.com/kenyon-wong/docs/blob/757fb85d879026c7d30eb19fafcf4cec231d8616/%E5%85%88%E7%9F%A5%E7%A4%BE%E5%8C%BA/JA-%E6%8C%87%E7%BA%B9%E8%AF%86%E5%85%A8%E7%B3%BB%E8%AE%B2%E8%A7%A3-%E4%B8%8A-%E5%85%88%E7%9F%A5%E7%A4%BE%E5%8C%BA/JA-%E6%8C%87%E7%BA%B9%E8%AF%86%E5%85%A8%E7%B3%BB%E8%AE%B2%E8%A7%A3-%E4%B8%8A-%E5%85%88%E7%9F%A5%E7%A4%BE%E5%8C%BA.md)
- [JA 指纹识全系讲解（下）](https://web.archive.org/web/20240422055319/https://xz.aliyun.com/t/14054?time__1311=mqmx9DBG0QD%3DNGNDQiiQGkfbOuiCdDcWoD)

- https://github.com/wwhtrbbtt/TrackMe
- https://en.wikipedia.org/wiki/Usage_share_of_web_browsers#Summary_tables
- https://www.fastly.com/blog/a-first-look-at-chromes-tls-clienthello-permutation-in-the-wild
- https://github.com/wi1dcard/fingerproxy

- https://github.com/google/boringssl/commit/e9c5d72c09e01a0f71f30f7c3454e5e7f8711476
- https://github.com/chromium/chromium/commit/08631bdfddaad0f25c62261734171674a9621484
- https://github.com/chromium/chromium/commit/8249eb7a1d2118bf9a6998c11964bae4c5db8b10
- https://github.com/chromium/chromium/commit/4493a1eb4595194a262617589c5a265de40e203e

- https://tlsfingerprint.io/
- https://github.com/net4people/bbs/issues/220

- https://github.com/refraction-networking/utls/blob/8f010b39328f36ec70d47a6f41c415d1f486aaff/u_parrots.go#L519-L589

- https://github.com/wangluozhe/requests
  - https://github.com/wangluozhe/chttp
- https://github.com/imroc/req
- https://github.com/gospider007/requests
  - https://github.com/gospider007/net/tree/4a6a7a20f9173a98d3142ab06d17f7dbd0e26a30
- https://medium.com/cu-cyber/impersonating-ja3-fingerprints-b9f555880e42
  - https://github.com/CUCyber/ja3transport
- https://github.com/refraction-networking/utls/issues/103
- https://github.com/fedosgad/mirror_proxy/
- https://github.com/refraction-networking/utls/pull/74
  - https://github.com/bassosimone/utlstransport
- https://github.com/saucesteals/mimic

- https://chromestatus.com/feature/5124606246518784

- https://github.com/rustls/rustls/issues/1125
- https://github.com/rustls/rustls/pull/1190
- https://github.com/rustls/rustls/issues/1421
- https://github.com/rustls/rustls/issues/1501
