#!/usr/bin/env bats

@test "Collections - Help is available" {
    ./id4i help collections | grep "Retrieve header data of the collection identified by --id"
    ./id4i help collections | grep "Retrieve elements (contained GUIDs) of the collection identified by --id"
}
