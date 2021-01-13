#!/bin/bash

peer chaincode invoke -o orderer0:7050 \
     --tls true --cafile $ORDERER_CA -C mychannel -n faulttol \
     --peerAddresses peer0-org1:7051 \
     --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1/peers/peer0-org1/tls/ca.crt \
     --peerAddresses peer0-org2:7051 \
     --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2/peers/peer0-org2/tls/ca.crt \
     -c '{"Args":["UpdateDeviceUsedResources", "__device__", "__ipc__" ,"__mem__","__bw__"]}' --waitForEvent
 