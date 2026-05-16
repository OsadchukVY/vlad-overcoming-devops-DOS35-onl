#!/bin/bash

set -o errexit
set -o pipefail
echo $#
if (( $# != 2 )); then
    echo "script <directory> <pattern>"
    exit 1
fi

DIR=$1
PAT=$2

if [[ ! -d "${DIR}" ]]; then
    echo "${DIR}" is not directory
    exit 2
fi

grep -R -l "${PAT}" "${DIR}" 2>/dev/null | while read -r file; do
    ls -lh "${file}"
done