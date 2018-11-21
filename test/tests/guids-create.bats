#!/usr/bin/env bats

@test "GUIDs - Create one GUID - short opts" {
    result=$(./id4i guids create -l 8 -c 1)
    [ $(echo $result | jq ".id4ns | length") -eq 1 ]
    [ $(echo $result | jq ".id4ns[0] | length") -eq 8 ]
}

@test "GUIDs - Create one GUID - long opts" {
    result=$(./id4i guids create --length 9 --count 1)
    [ $(echo $result | jq ".id4ns | length") -eq 1 ]
    [ $(echo $result | jq ".id4ns[0] | length") -eq 9 ]
}

@test "GUIDs - Create multiple GUIDs" {
    result=$(./id4i guids create -l 15 -c 10)
    [ $(echo $result | jq ".id4ns[0] | length") -eq 15 ]
    [ $(echo $result | jq ".id4ns[3] | length") -eq 15 ]
    [ $(echo $result | jq ".id4ns[9] | length") -eq 15 ]
    [ $(echo $result | jq ".id4ns | length") -eq 10 ]

    [ 1 -eq 2 ]
}
