#!/usr/bin/env bash

set -o errexit  # exit on error
set -o nounset  # don't allow unset variables
# set -o xtrace # enable for debugging

go run main.go help > doc/help.txt
go run main.go help guids > doc/help-guids.txt
go run main.go help storage > doc/help-storage.txt
go run main.go help history > doc/help-history.txt
