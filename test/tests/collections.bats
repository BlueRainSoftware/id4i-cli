#!/usr/bin/env bats

setup() {
    NOW=$(date +%s)
}

@test "Collections - Help is available" {
    ./id4i help collections | grep "Retrieve header data of the collection identified by --id"
    ./id4i help collections | grep "Retrieve elements (contained GUIDs) of the collection identified by --id"
}

@test "Collections - Create labelled collection - short opts" {
    collectionId=$(./id4i collections create -t LABELLED_COLLECTION -a "Logistic collection $NOW" | jq -r ".id4n")
    [ $(echo $collectionId | wc -m) -eq 11 ]
    ./id4i collections delete --id ${collectionId}
}

@test "Collections - Create labelled collection - long opts" {
    collectionId=$(./id4i collections create --length 15 --type LABELLED_COLLECTION --label "Labelled collection $NOW" | jq -r ".id4n")
    [ $(echo $collectionId | wc -m) -eq 16 ]
    ./id4i collections delete --id ${collectionId}
}

@test "Collections - Create logistic collection" {
    collectionId=$(./id4i collections create --length 35 --type LOGISTIC_COLLECTION --label "Logistic collection 2 $NOW" | jq -r ".id4n")
    [ $(echo $collectionId | wc -m) -eq 36 ]
    ./id4i collections delete -i ${collectionId}
}

@test "Collections - Create routing collection" {
    collectionId=$(./id4i collections create --length 101 --type ROUTING_COLLECTION --label "routing collection $NOW" | jq -r ".id4n")
    [ $(echo $collectionId | wc -m) -eq 102 ]
    ./id4i collections delete -i ${collectionId}
}
