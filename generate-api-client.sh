#!/usr/bin/env bash

set -o errexit  # exit on error
set -o nounset  # don't allow unset variables
# set -o xtrace # enable for debugging

usage() {
  printf "Generate ID4i API client library for golang\n"

  printf "Usage: $(basename "$0") "
  printf -- "[-h] "
  printf -- "[-v] "
  printf -- "[-c] "
  printf -- "[-d] "
  printf "\n"

  printf "  -%s\t%s - %s%s\n" "h" "help" "Show this help message." ""
  printf "  -%s\t%s - %s%s\n" "v" "version" "Show version information." ""
  printf "  -%s\t%s - %s%s\n" "c" "current" "Pull current swagger.json from production" " (default: 0)"
  printf "  -%s\t%s - %s%s\n" "d" "deps" "Install go dependencies" " (default: 0)"
}

version() {
  printf "0.0.3\n"
}

# default values
current=false
deps=false

# option parsing
OPTSPEC=:hvcd
while getopts $OPTSPEC option; do
  case "$option" in
    h ) usage; exit 0  ;;
    v ) version; exit 0  ;;
    c )  current=true;;
    d )  deps=true;;
   \? ) echo "Unknown option: -$OPTARG" >&2; exit 1;;
    : ) echo "Missing option argument for -$OPTARG" >&2; exit 1;;
    * ) echo "Unimplemented option: -$OPTARG" >&2; exit 1;;
  esac
done
shift $((OPTIND - 1))

# convenience variables
__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__file="${__dir}/$(basename "${BASH_SOURCE[0]}")"
__base="$(basename ${__file} .sh)"
__root="$(cd "$(dirname "${__dir}")" && pwd)" # update this to make it point your project's root

if [ "$current" = true ] ; then
    rm -f swagger.json
    wget https://backend.id4i.de/docs/swagger.json
fi

if [ "$deps" = true ] ; then
    go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
fi

rm -rf api/
swagger generate client --additional-initialism ID4i --additional-initialism ID4n --additional-initialism GET -m api/models -c api/client -f swagger.json
