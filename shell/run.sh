#!/usr/bin/env bash
BASEDIR=$(dirname "$0")
TS=$(date '+%d%m_%H-%M');

nohup /usr/local/go/bin/go run "$BASEDIR"/../main.go > "$BASEDIR"/../logs/"$TS".log &