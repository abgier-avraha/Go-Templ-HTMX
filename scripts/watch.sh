#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
MD5_BIN="md5sum"

if [[ "$OSTYPE" == "darwin"* ]]; then
  MD5_BIN="md5"
fi

# Bug where nodemon can't watch parent folders
cd $SCRIPT_DIR/..

WATCH="./"
ACTION=./scripts/run.sh

nodemon --watch $WATCH -e templ,.go --signal SIGTERM --exec $ACTION
