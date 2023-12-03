#!/usr/bin/env bash
set -e

go build -ldflags "-s -w"
echo "Before compressing binary:"
ls -l clexpg
upx clexpg
echo "After compressing binary:"
ls -l clexpg
