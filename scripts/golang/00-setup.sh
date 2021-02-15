#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

checkmark="(\xE2\x9C\x94)"
repo_name=""
mod_file="./go.mod"

if [ ! -f "$mod_file" ]; then
    read -p "What is your Go module name ? <user/repo>: " repo_name
    go mod init $repo_name
    echo -e "1. Init Go module $checkmark"
else
    echo -e "1. Go module exists $checkmark"
fi

mkdir -p {cmd,internal,$repo_name}
echo "ENV=dev" > ./.env.example
cp ./.env.example .env
echo -e "2. Init Go scaffold $checkmark"
