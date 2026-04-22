#!/usr/bin/env bash
set -euo pipefail

read -r VERSION <VERSION
tags=()

for dir in */; do
  tag="${dir}${VERSION}"
  git tag "$tag"
  tags+=("$tag")
done

git push origin "${tags[@]}"
