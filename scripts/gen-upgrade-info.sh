#!/usr/bin/env bash

#  get version as first parameter
VERSION=$1

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

GITHUB_REPO=tokenize-titan/titan

# verify release version exists on github

# Get the latest release
release_url="https://api.github.com/repos/$GITHUB_REPO/releases/tags/$VERSION"
release_info=$(curl -s $release_url)

# Check if the release is not found
if echo $release_info | jq '.message' | grep -q "Not Found"; then
    echo "ERROR: release $VERSION not found"
    exit 1
fi

# Extract asset checksums
asset_checksums=$(echo $release_info | jq -r '.assets[] | select(.name=="checksums.txt") | .browser_download_url')

# check if checksums.txt is not exists in assets
if [ -z "$asset_checksums" ]; then
    echo "ERROR: checksums.txt not found in release $VERSION"
    exit 1
fi

# get checksums info
checksums=$(curl -sL $asset_checksums)

# get checksums for each file
# checksums have format: <checksum>  <filename>  

# get list of files
files=$(echo "$checksums" | awk '{print $2}')

# create upgrade info in format json:
# {
#   "binaries": {
#     "darwin/arm64": "<download_url>?checksum=<checksum>",
#   }
# }

upgrade_info='{"binaries":{'
for file in $files; do
  # get checksum for file
  checksum=$(echo "$checksums" | grep $file | awk '{print $1}')  
  # get file name without extension
  filename="${file%.tar.gz}"
  # get os and arch
  os_arch="${filename#*_*_}"
  # replace _ with / in os_arch
  os_arch="${os_arch//_//}"
  # to lowercase
  os_arch=$(echo $os_arch | tr '[:upper:]' '[:lower:]')
  # get download url
  download_url=$(echo $release_info | jq -r ".assets[] | select(.name==\"$file\") | .browser_download_url")
  # add to upgrade info
  upgrade_info="$upgrade_info\"$os_arch\":\"$download_url?checksum=$checksum\","
done

# remove last comma
upgrade_info=$(echo $upgrade_info | sed 's/,$//')
upgrade_info="$upgrade_info}}"

echo $upgrade_info

