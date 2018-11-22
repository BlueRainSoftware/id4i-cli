#!/usr/bin/env bats

setup() {
    source .preflightData
}

@test "Basic - Help is available" {
    ./id4i help | grep 'ID4i API commandline application.'
    ./id4i help | grep 'id4i \[command\]'
    ./id4i help | grep -- '--apikey string         ID4i API key to use'
    ./id4i help | grep -- '-o, --organization string   ID4i organization namespace to work in'
}


@test "Basic - Default config file is used" {
    ./id4i info --show-config 2>&1 | grep ${PWD}/.id4i.properties
    ./id4i info --show-config 2>&1 | grep ${APIKEY_ID}
}
