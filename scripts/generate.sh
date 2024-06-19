#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Generate templates
templ generate

# Generate tailwind CSS
$SCRIPT_DIR/../node_modules/.bin/tailwindcss -i ./public/tailwind.css -o ./public/tailwind.generated.css

# Compeile Go
go build $SCRIPT_DIR/../