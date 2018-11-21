#!/usr/bin/env bats

setup() {
    source .preflightData
    guid=$(./id4i guids create -l 8 -c 1 | jq -r ".id4ns[0]")
}

@test "History - Add history item" {
    ./id4i history add -i $guid --type PRODUCTION_STARTED
    sleep 1
    [ $(./id4i history list -i $guid --include-private | jq ".elements | length") -eq "1" ]
}

@test "History - Add public history item" {
    ./id4i history add -i $guid --type CREATED
    ./id4i history add -i $guid --type PRODUCTION_STARTED --publish
    sleep 1
    [ $(./id4i history list -i $guid | jq ".elements | length") -eq "1" ]
    [ $(./id4i history list -i $guid --include-private | jq ".elements | length") -eq "2" ]
}

@test "History - Add multiple history items" {
    ./id4i history add -i $guid --type CREATED
    ./id4i history add -i $guid --type PRODUCTION_STARTED
    ./id4i history add -i $guid --type PRODUCTION_CANCELLED
    ./id4i history add -i $guid --type RECYCLED
    ./id4i history add -i $guid --type CREATED
    sleep 1
    [ $(./id4i history list -i $guid --include-private | jq ".elements | length") -eq "5" ]
}

@test "History - Share history item" {
    sleep 1
    [ $(./id4i history list -i $guid --include-private | jq ".elements[0].visibility.sharedOrganizationIds | length") -eq "0" ]
    ./id4i history set-visibility -i $guid -s de.id4i.company1 --sequence 0
    sleep 1
    [ $(./id4i history list -i $guid --include-private | jq ".elements[0].visibility.sharedOrganizationIds | length") -eq "1" ]
}

@test "History - Publish history item" {
    sleep 1
    [ $(./id4i history list -i $guid | jq ".elements | length") -eq "0" ]
    [ $(./id4i history list -i $guid -p | jq ".elements | length") -eq "1" ]
    ./id4i history set-visibility -i $guid -s de.id4i.company1 -p --sequence 0
    sleep 1
    [ $(./id4i history list -i $guid | jq ".elements | length") -eq "1" ]
}
