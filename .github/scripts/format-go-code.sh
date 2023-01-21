#!/usr/bin/env bash

set -eu

FILEPATH="$1"

gofmt -s -w "$FILEPATH"

# https://github.com/rinchsan/gosimports
gosimports -local github.com/ydb-platform/terraform-provider-ydb -w "$FILEPATH"

# https://github.com/mvdan/gofumpt
gofumpt -w "$FILEPATH"
