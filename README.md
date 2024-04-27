# anti fingerprinting

This project aims at anti-fingerprinting **programming**, it consists of the following components:

- [ ] [**Collector**](./cmd/collector)

  > It collects and records fingerprints from requests, currently it's based on [TrackMe](https://github.com/wwhtrbbtt/TrackMe), [fingerproxy](https://github.com/wi1dcard/fingerproxy), and [clienthellod](https://github.com/gaukas/clienthellod).

  - [x] HTTP/2
  - [x] TLS
  - [ ] QUIC
  - [ ] TCP/IP

- [ ] **Automation**

  > It sends requests to **Collector** automatically from the most popular platforms and browsers, with the help of container/vm/emulator or other technologies.  
  > The collected fingerprints can be found in [the assets branch](https://github.com/hellodword/anti-fingerprinting/tree/assets).

  - [x] Windows (win10x64 and win11x64 running in [dockur/windows](https://github.com/dockur/windows))

    - [x] Chrome
    - [x] Firefox
    - [x] Edge

  - macOS

    - [ ] Safari

  - Android

    - [ ] Webview
    - [ ] Chrome
    - [ ] Sansung Internet
    - [ ] UC

  - iOS

    - [ ] Safari

- [x] **Modifier**

  - [x] [TLS](https://github.com/refraction-networking/utls)
  - [x] [HTTP/2](https://github.com/hellodword/http2-custom-fingerprint)
  - [x] [QUIC](https://github.com/refraction-networking/uquic)

- [ ] **Generator**

  > It generate code for **Modifier** from the collected fingerprints, or provide it as a library for dynamic usage.

  - [ ] HTTP/2
  - [ ] TLS
  - [ ] QUIC

---

## Reference

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
