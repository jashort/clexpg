#!/usr/bin/env bash
set -e

go build -ldflags "-s -w"
echo "Before compressing binary:"
ls -l clexpg
# UPX no longer supported on MacOS:
# upx: clexpg: CantPackException: macOS is currently not supported (try --force-macos)
#upx clexpg
#echo "After compressing binary:"
#ls -l clexpg
