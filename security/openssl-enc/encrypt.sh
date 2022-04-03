#!/usr/bin/env bash

# standard bash error handling
set -o nounset # treat unset variables as an error and exit immediately.
set -o errexit # exit immediately when a command fails.

CURRENT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
readonly CURRENT_DIR

encrypt() {
    local input_file=$1
    local output_file=$2
    openssl enc -in "${input_file}" -aes-256-cbc -out "${output_file}"
}

encrypt "${CURRENT_DIR}/assets/decrypted/example.yaml" "${CURRENT_DIR}/example.yaml.enc"

echo "Files encrypted successfully"
