#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Bug where nodemon can't watch parent folders
cd $SCRIPT_DIR/..

WATCH="./"
ACTION=./scripts/run.sh

./node_modules/.bin/nodemon --watch $WATCH -e templ,.go --signal SIGTERM --exec $ACTION