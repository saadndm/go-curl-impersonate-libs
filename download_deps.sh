#!/usr/bin/env bash
set -euo pipefail

read -r VERSION <VERSION
URL=https://github.com/saadndm/curl-impersonate

for dir in */; do
  platform=${dir%/}
  file="libcurl-impersonate-${VERSION}.${platform}"
  curl -fLo "${dir}${file}.tar.gz" "${URL}/releases/download/${VERSION}/${file}.tar.gz"
  tar -xzf "${dir}${file}.tar.gz" -C "$dir"
  rm -f "${dir}${file}.tar.gz"
  find "$dir" -maxdepth 1 -type f \
    ! -name 'go.mod' \
    ! -name 'link.go' \
    ! -name 'smoke.go' \
    ! -name 'smoke_test.go' \
    ! -name 'libcurl-impersonate.a' \
    -delete
  rm -rf "${dir}bin" "${dir}include" "${dir}lib"
done
