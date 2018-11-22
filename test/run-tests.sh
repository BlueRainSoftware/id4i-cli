#!/usr/bin/env bash

set -o errexit  # exit on error
set -o nounset  # don't allow unset variables
# set -o xtrace # enable for debugging

usage() {
  printf "Run ID4i CLI tests\n"

  printf "Usage: $(basename "$0") "
  printf -- "[-h] "
  printf -- "[-v] "
  printf -- "[-i] "
  printf -- "[-p] "
  printf -- "[-c] "
  printf -- "[-b] "
  printf -- "[-t] "
  printf "\n"

  printf "  -%s\t%s - %s%s\n" "h" "help" "Show this help message." ""
  printf "  -%s\t%s - %s%s\n" "v" "version" "Show version information." ""
  printf "  -%s\t%s - %s%s\n" "i" "install" "Install Prerequisites" ""
  printf "  -%s\t%s - %s%s\n" "p" "preflight" "Run preflight script (provision ID4i test user)" ""
  printf "  -%s\t%s - %s%s\n" "c" "cleanup" "Clean up test results after successful tests" ""
  printf "  -%s\t%s - %s%s\n" "b" "build" "Build ID4i binary before testing (requires Go)" ""
  printf "  -%s\t%s - %s%s\n" "t" "tap-format" "Use TAP format test output (for CI reports)" ""
}

version() {
  printf "0.0.1\n"
}

# default values
opt_install=""
opt_preflight=""
opt_cleanup=""
opt_build=""
opt_tap=""

# option parsing
OPTSPEC=:hvipcbt
while getopts $OPTSPEC option; do
  case "$option" in
    h ) usage; exit 0  ;;
    v ) version; exit 0  ;;
    i ) opt_install=1;  ;;
    p ) opt_preflight=1;  ;;
    c ) opt_cleanup=1;  ;;
    b ) opt_build=1;  ;;
    t ) opt_tap="--tap";  ;;
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

# this would be a good place to start writing your actual script.

if [ "$opt_install" = 1 ] ; then
    echo "Install testing tools"
    ./install-test-tools.sh
fi

if [ "$opt_build" = 1 ] ; then
    echo "Build binary for tests"
    go build -o id4i ../main.go
fi

if [ "$opt_preflight" = 1 ] ; then
    echo "Provision first test user in ID4i"
    ./preflight.sh
    echo "Prepare id4i cli configuration"
    source .preflightData
    echo "backend=id4i-develop.herokuapp.com" > ./.id4i.properties
    echo "apikey=$APIKEY_ID" >> ./.id4i.properties
    echo "secret=$PASSWORD" >> ./.id4i.properties
    echo "organization=$ORGANIZATION" >> ./.id4i.properties
    mv ./.preflightData ./.preflightData.1

    echo "Provision second test user in ID4i (for transfers)"
    ./preflight.sh
    echo "Prepare id4i cli configuration"
    source .preflightData
    echo "backend=id4i-develop.herokuapp.com" > ./.id4i.2.properties
    echo "apikey=$APIKEY_ID" >> ./.id4i.2.properties
    echo "secret=$PASSWORD" >> ./.id4i.2.properties
    echo "organization=$ORGANIZATION" >> ./.id4i.2.properties
    mv ./.preflightData ./.preflightData.2

    mv ./.preflightData.1 ./.preflightData
fi

echo "Run tests"
./bats/bin/bats $opt_tap tests/*.bats | tee tests.log

if [ "$opt_cleanup" = 1 ] ; then
    echo "Cleaning up"
    rm id4i
    rm .id4i.properties
    rm .id4i.2.properties
    rm .preflightData
    rm .preflightData.2
    rm -rf bats
    echo "Clean up finished"
fi

echo "All done"
