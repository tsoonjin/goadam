#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

ROOT_PATH=$1
SECRET_PATH=$2
IFS='='

echo "Preparing secrets"
[ ! -f $SECRET_PATH ] && { echo "$SECRET_PATH file not found"; exit 99; }
while read id value
do
	echo $value >| $ROOT_PATH/secrets/$id.txt
done < $SECRET_PATH
