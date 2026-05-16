#!/bin/bash

set -o errexit
set -o pipefail

if (( $@ != 1 )); then
    echo "Pass directory"
    exit 1
fi

FILES=$(find "$1" 2>/dev/null)

for i in $FILES; do
    ls -lh $i
done