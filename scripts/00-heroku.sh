#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

# Variable declarations
checkmark="(\xE2\x9C\x94)"
repo_name=""
project_name=""

if ! command -v heroku &> /dev/null
then
    echo -n "Heroku client not found. Do you want to automatically install ? (Y/n)"
    read answer
    if [ "$answer" != "${answer#[Yy]}" ] ;then
        echo -n "Installing heroku client ..."
        curl https://cli-assets.heroku.com/install.sh | sh
    else
        echo -n "Heroku client not found. Failed to proceed with setup"
        exit
    fi
fi

echo -e "Heroku Client: heroku installed $checkmark\n"
echo -e "1. Setup heroku login"
heroku container:login
