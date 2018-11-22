#!/usr/bin/env bats

setup() {
    guid=$(./id4i guids create -l 8 -c 1 | jq -r ".id4ns[0]")
}

@test "Transfer - Help is available" {
    ./id4i help transfer | grep "send        Prepare the transfer of an ID to another organization"
}

@test "Transfer - Transfer GUID" {
}

@test "Transfer - Transfer GUID keeping ownership" {
}

@test "Transfer - Transfer to multiple recipients" {
}

@test "Transfer - Open GUID for claims" {
}

