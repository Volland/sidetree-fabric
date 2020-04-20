#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@blockchain-handler
Feature:
  Background: Setup
    Given DCAS collection config "dcas-cfg" is defined for collection "dcas" as policy="OR('Org1MSP.member','Org2MSP.member')", requiredPeerCount=1, maxPeerCount=2, and timeToLive=
    Given off-ledger collection config "diddoc-cfg" is defined for collection "diddoc" as policy="OR('IMPLICIT-ORG.member')", requiredPeerCount=0, maxPeerCount=0, and timeToLive=
    Given off-ledger collection config "fileidx-cfg" is defined for collection "fileidxdoc" as policy="OR('IMPLICIT-ORG.member')", requiredPeerCount=0, maxPeerCount=0, and timeToLive=
    Given off-ledger collection config "meta-data-cfg" is defined for collection "meta_data" as policy="OR('IMPLICIT-ORG.member')", requiredPeerCount=0, maxPeerCount=0, and timeToLive=

    Given the channel "mychannel" is created and all peers have joined

    # Give the peers some time to gossip their new channel membership
    And we wait 20 seconds

    And "system" chaincode "configscc" is instantiated from path "in-process" on the "mychannel" channel with args "" with endorsement policy "AND('Org1MSP.member','Org2MSP.member')" with collection policy ""
    And "system" chaincode "sidetreetxn" is instantiated from path "in-process" on the "mychannel" channel with args "" with endorsement policy "AND('Org1MSP.member','Org2MSP.member')" with collection policy "dcas-cfg"
    And "system" chaincode "document" is instantiated from path "in-process" on the "mychannel" channel with args "" with endorsement policy "OR('Org1MSP.member','Org2MSP.member')" with collection policy "diddoc-cfg,fileidx-cfg,meta-data-cfg"

    And fabric-cli network is initialized
    And fabric-cli plugin "../../.build/ledgerconfig" is installed
    And fabric-cli context "mychannel" is defined on channel "mychannel" with org "peerorg1", peers "peer0.org1.example.com,peer1.org1.example.com,peer2.org1.example.com" and user "User1"

    And we wait 10 seconds

    Then fabric-cli context "mychannel" is used
    And fabric-cli is executed with args "ledgerconfig update --configfile ./fixtures/config/fabric/mychannel-consortium-config.json --noprompt"
    And fabric-cli is executed with args "ledgerconfig update --configfile ./fixtures/config/fabric/mychannel-org1-config.json --noprompt"
    And fabric-cli is executed with args "ledgerconfig update --configfile ./fixtures/config/fabric/mychannel-org2-config.json --noprompt"

    # Wait for the Sidetree services to start up on mychannel
    And we wait 10 seconds

  @blockchain_s1
  Scenario: Blockchain functions
    When an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/version"
    Then the JSON path "name" of the response equals "Hyperledger Fabric"
    And the JSON path "version" of the response equals "2.0.0"

    When an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/time"
    Then the JSON path "time" of the response is not empty
    And the JSON path "hash" of the response is not empty
    And the JSON path "time" of the response is saved to variable "time"
    And the JSON path "hash" of the response is saved to variable "hash"

    When an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/time/${hash}"
    Then the JSON path "hash" of the response equals "${hash}"
    And the JSON path "time" of the response equals "${time}"

    # Invalid hash - Bad Request (400)
    Then an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/time/xxx_xxx" and the returned status code is 400

    # Hash not found - Not Found (404)
    Then an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/time/AQIDBAUGBwgJCgsM" and the returned status code is 404

    # Write a few Sidetree transactions. Scatter the requests across different endpoints to generate multiple
    # Sidetree transactions within the same block. The Orderer's batch timeout is set to 2s, so sleep 2s between
    # writes to guarantee that we generate a few blocks.
    Then client sends request to "https://localhost:48326/document" to create DID document in namespace "did:sidetree"
    And check success response contains "#didDocumentHash"
    And client sends request to "https://localhost:48327/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48328/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48426/document" to create DID document in namespace "did:sidetree"

    Then we wait 2 seconds

    Then client sends request to "https://localhost:48427/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48428/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48326/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48327/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48328/document" to create DID document in namespace "did:sidetree"

    Then we wait 2 seconds

    Then client sends request to "https://localhost:48426/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48427/document" to create DID document in namespace "did:sidetree"
    And client sends request to "https://localhost:48428/document" to create DID document in namespace "did:sidetree"

    Then we wait 15 seconds

    # The config setting for maxTransactionsInResponse is 10 so we should expect 10 transactions in the query for all transactions
    When an HTTP GET is sent to "https://localhost:48326/sidetree/0.1.3/blockchain/transactions"
    And the JSON path "moreTransactions" of the boolean response equals "true"
    And the JSON path "transactions.9.transactionTimeHash" of the response is not empty
    And the JSON path "transactions.9.anchorString" of the response is not empty