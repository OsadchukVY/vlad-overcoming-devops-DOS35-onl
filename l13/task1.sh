#!/bin/bash

set -o errexit
set -o pipefail

if (( $# != 1 )); then
    echo "script <directory> "
    exit 1
fi

DIR=$1

if [[ ! -d "${DIR}" ]]; then
    echo "${DIR}" is not directory
    exit 2
fi


find "$1" -type f -name "*.log" 2>/dev/null | while read -r file; do
    cp "${file}" "${file}"_"$(date +%s)".bak
done


find "$1" -type f -name "*.py" 2>/dev/null | while read -r file; do
    cp "${file}" "${file}"_"$(git rev-parse --short HEAD)".bak
done