#!/usr/bin/env bats

setup() {
    guid=$(./id4i guids create -l 8 -c 1 | jq -r ".id4ns[0]")
}

@test "Storage - Help is available" {
    ./id4i help storage | grep "upload      Upload new document"
}

@test "Storage - Upload document" {
    ./id4i storage upload -f test-user-names.1.txt -i ${guid} | grep "filename\":\"test-user-names.1.txt\""
}

@test "Storage - Upload and rename document" {
    ./id4i storage upload -f test-user-names.1.txt -d myNewDocument.md -i ${guid} | grep "filename\":\"myNewDocument.md\""
}

@test "Storage - Default visibility of document is private unshared" {
    ./id4i storage upload -f test-user-names.1.txt -i ${guid} | grep -v "public\":true"
    ./id4i storage upload -f test-user-names.2.txt -i ${guid} | grep "sharedOrganizationIds\":\[\]"
}

@test "Storage - Publish document" {
    ./id4i storage upload -f test-user-names.1.txt -i ${guid} -p | grep "public\":true"
}

@test "Storage - Share document" {
    ./id4i storage upload -f test-user-names.1.txt -i ${guid} --share-with de.id4i.company1 | grep -v "public\":true"
    ./id4i storage upload -f test-user-names.2.txt -i ${guid} --share-with de.id4i.company1 | grep "sharedOrganizationIds\":\[\"de.id4i.company1\"\]"

}
