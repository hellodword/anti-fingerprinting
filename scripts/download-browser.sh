#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./windows/oem/

[ -n "$BROWSER_URL" ]
[ -n "$BROWSER_BROWSER" ]
[ -n "$BROWSER_VERSION" ]

version_path="./windows/oem/$BROWSER_BROWSER-$BROWSER_VERSION"
[ -f "$version_path" ] || \
aria2c --continue --out "$version_path" "$BROWSER_URL"

[ -z "$BROWSER_HASH" ] || \
[ "$(sha256sum "$version_path" | awk '{print $1}')" = "$BROWSER_HASH" ] || \
[ "$(sha1sum "$version_path" | awk '{print $1}')" = "$BROWSER_HASH" ] || \
[ "$(sha512sum "$version_path" | awk '{print "sha512:" $1}')" = "$BROWSER_HASH" ] || \
(rm -rf "$version_path"; exit 1)

if [[ "$BROWSER_URL" == *.cab ]]; then
  which cabextract || sudo apt-get install -y cabextract
  cabextract "$version_path"
  mv MicrosoftEdgeEnterpriseX64.msi ./windows/oem/${BROWSER_BROWSER}_installer.exe
else
  mv "$version_path" ./windows/oem/${BROWSER_BROWSER}_installer.exe
fi
