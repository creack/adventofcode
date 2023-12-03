#!/usr/bin/env sh

year=$1
day=$2

curl 'https://adventofcode.com/'"${year}"'/day/'"${day}"'/input' \
     -H 'cookie: session='"${COOKIE}"
