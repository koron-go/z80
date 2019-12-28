#!/bin/sh

set -eu

name=$1 ; shift

grep -lr '{{\.Name}}' . | xargs sed -i.bak -e "s/{{\.Name}}/${name}/g"
find . -type f -name \*.bak | xargs rm

#rm -f $0
