#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

# Variable declarations
checkmark="(\xE2\x9C\x94)"

if ! command -v gh &> /dev/null
then
    echo -n "Github client not found. Do you want to automatically install ? (Y/n)"
    read answer
    if [ "$answer" != "${answer#[Yy]}" ] ;then
        echo -n "Installing github client ..."
        sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0
        sudo apt-add-repository https://cli.github.com/packages
        sudo apt update
        sudo apt install gh
    else
        echo -n "Github client not found. Failed to proceed with setup"
        exit
    fi
fi

echo -e "Github Client: gh installed $checkmark\n"
echo -e "1. Create New Github Repo"
gh repo create
