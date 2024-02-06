#!/usr/bin/env bash

PACKAGE_NAME=github.com/tokenize-titan/titan
GOLANG_CROSS_VERSION=v1.20

# verify if `MAKE_PROJECT_ROOT` and `MAKE_BUILD_TAGS` is set
if [ -z "$MAKE_PROJECT_ROOT" ] || [ -z "$MAKE_BUILD_TAGS" ]; then
  echo "This script should be called from Makefile. Use 'make release-dry-run' or 'make release' instead."
  exit 1
fi

PROJECT_ROOT=$MAKE_PROJECT_ROOT
BUILD_TAGS=$MAKE_BUILD_TAGS
BUILD_TAGS_COMMA_SEP=$MAKE_BUILD_TAGS_COMMA_SEP

GOPATH=$(go env GOPATH)
# get project root directory
WASM_PATH=$(go list -f '{{.Dir}}' github.com/CosmWasm/wasmvm)

echo VERSION=$VERSION
echo TMVERSION=$TMVERSION
echo PROJECT_ROOT=$PROJECT_ROOT
echo BUILD_TAGS=$BUILD_TAGS

dry_run=0
for arg in "$@"
do
    if [ "$arg" = "--dry-run" ]; then
        dry_run=1
        break
    fi
done

# remove tmp directory
rm -rf tmp
# copy lib file to `dist/lib`
mkdir -p tmp/lib
cp $WASM_PATH/internal/api/*.so tmp/lib
cp $WASM_PATH/internal/api/*.dylib tmp/lib



# if `--dry-run` is set then skip publishing
if [ "$dry_run" -eq 1 ]; then
  docker run \
      --rm \
      --privileged \
      -e TMVERSION=$TMVERSION \
      -e VERSION=$VERSION \
      -e BUILD_TAGS="$BUILD_TAGS" \
      -e BUILD_TAGS_COMMA_SEP="$BUILD_TAGS_COMMA_SEP" \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v $PROJECT_ROOT:/go/src/$PACKAGE_NAME \
      -v ${GOPATH}/pkg:/go/pkg \
      -w /go/src/$PACKAGE_NAME ghcr.io/goreleaser/goreleaser-cross:$GOLANG_CROSS_VERSION \
      --clean --skip-validate --skip-publish --snapshot

else
  # check file `.release-env` exists
  if [ ! -f .release-env ]; then
    echo "File .release-env not found"
    exit 1
  fi

  docker run \
		--rm \
		--privileged \
    -e TMVERSION=$TMVERSION \
    -e VERSION=$VERSION \
    -e BUILD_TAGS="$BUILD_TAGS" \
    -e BUILD_TAGS_COMMA_SEP="$BUILD_TAGS_COMMA_SEP" \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $PROJECT_ROOT:/go/src/$PACKAGE_NAME \
		-w /go/src/$PACKAGE_NAME ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--clean --skip-validate
fi

# remove tmp directory
rm -rf tmp