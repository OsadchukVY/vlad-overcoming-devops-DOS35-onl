#!/bin/bash

set -o errexit
set -o pipefail

if (( $# != 1 )); then
    echo "script <file>"
    exit 1
fi

DIR=$1
PAT="([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\"([]!#-[^-~ \t]|(\\[\t -~]))+\")@([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\[[\t -Z^-~]*])"

if [[ ! -f "${DIR}" ]]; then
    echo "${DIR}" is not directory
    exit 2
fi

echo "$(grep -Eiorh "${PAT}" "${DIR}" 2>/dev/null)"
