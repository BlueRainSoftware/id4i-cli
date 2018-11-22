#!/usr/bin/env bats

setup() {
    guid=$(./id4i guids create -l 8 -c 1 | jq -r ".id4ns[0]")

    source .preflightData

    # Put all data from the second user into environment variables
    IFS=$'\n'
    for line in $( cat .preflightData.2 ); do export U2_$line; done;
}

@test "Transfer - Help is available" {
    ./id4i help transfer | grep "send        Prepare the transfer of an ID to another organization"
}

@test "Transfer - Transfer GUID" {
    ./id4i transfer status -i ${guid} | grep "recipientOrganizationIds\":\[\]"
    ./id4i transfer send -i ${guid} -r ${U2_ORGANIZATION}
    ./id4i transfer status -i ${guid} | grep "recipientOrganizationIds\":\[\"${U2_ORGANIZATION}\"\]"
    ./id4i transfer status -i ${guid} | grep "keepOwnership\":false"
}

@test "Transfer - Transfer GUID keeping ownership" {
    ./id4i transfer status -i ${guid} | grep "recipientOrganizationIds\":\[\]"
    ./id4i transfer send -i ${guid} -k -r ${U2_ORGANIZATION}
    ./id4i transfer status -i ${guid} | grep "recipientOrganizationIds\":\[\"${U2_ORGANIZATION}\"\]"
    ./id4i transfer status -i ${guid} | grep "keepOwnership\":true"
}

@test "Transfer - Transfer to multiple recipients" {
    ./id4i transfer status -i ${guid} | grep "recipientOrganizationIds\":\[\]"
    ./id4i transfer send -i ${guid} -r ${U2_ORGANIZATION} -r ${ORGANIZATION}
    [ $(./id4i transfer status -i $guid | jq ".recipientOrganizationIds | length") -eq "2" ]
}

@test "Transfer - Open GUID for claims" {
    ./id4i transfer status -i ${guid} | grep "openForClaims\":false"
    ./id4i transfer send -i ${guid} -c
    [ $(./id4i transfer status -i $guid | jq ".recipientOrganizationIds | length") -eq "0" ]
    ./id4i transfer status -i ${guid} | grep "openForClaims\":true"
}

