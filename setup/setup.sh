#!/bin/bash

function generateCerts(){

    if [ ! -f $GOPATH/bin/cryptogen ]; then
        go get github.com/hyperledger/fabric/common/tools/cryptogen
    fi
    
    echo
        echo "##########################################################"
        echo "##### Generate certificates using cryptogen tool #########"
        echo "##########################################################"
        if [ -d ./crypto-config ]; then
                rm -rf ./crypto-config
        fi

    $GOPATH/bin/cryptogen generate --config=./crypto-config.yaml
    echo
}


function generateChannelArtifacts(){

    if [ ! -d ./channel-artifacts ]; then
                mkdir channel-artifacts
        fi

        if [ ! -f $GOPATH/bin/configtxgen ]; then
        go get github.com/hyperledger/fabric/common/tools/configtxgen
    fi

    echo
        echo "#################################################################"
        echo "### Generating channel configuration transaction 'channel.tx' ###"
        echo "#################################################################"

    $GOPATH/bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID testchainid -outputBlock ./channel-artifacts/genesis.block
    $GOPATH/bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID "mychannel"


    echo
        echo "#################################################################"
        echo "#######    Generating anchor peer update for Org1MSP   ##########"
        echo "#################################################################"
        $GOPATH/bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/org1Anchors.tx  -channelID "mychannel" -asOrg org1MSP

        echo
        echo "#################################################################"
        echo "#######    Generating anchor peer update for Org2MSP   ##########"
        echo "#################################################################"
        $GOPATH/bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/org2Anchors.tx  -channelID "mychannel" -asOrg org2MSP
        echo
}


# Install the binaries 
wget -N https://github.com/hyperledger/fabric/releases/download/v2.2.0/hyperledger-fabric-linux-amd64-2.2.0.tar.gz

tar -xzf hyperledger-fabric-linux-amd64-2.2.0.tar.gz

# Move to the bin path
mv bin/* /bin

# Check that you have successfully installed the tools by executing
configtxgen --version

generateCerts
generateChannelArtifacts

# You need to have an NFS created at edge nodes
scp -r crypto-config pi@192.168.1.233:/home_nfs/
scp -r channel-artifacts pi@192.168.1.233:/home_nfs/

# Create namespace
kubectl create ns hyperledger

# Create orderer
kubectl create -f orderer-service/

kubectl create -f org1/
kubectl create -f org2/

## Open channel cli and run the following

kubectl exec -ti cli_org1_pod_name sh -n hyperledger

peer channel create -o orderer0:7050 -c mychannel -f ./scripts/channel-artifacts/channel.tx --tls true --cafile $ORDERER_CA

peer channel join -b mychannel.block

## Open cli from org 2 to join the channel
kubectl exec -ti cli_org2_pod_name sh -n hyperledger

peer channel fetch 0 mychannel.block -c mychannel -o orderer0:7050 --tls --cafile $ORDERER_CA

peer channel join -b mychannel.block

peer channel list


# Installing the external chaincode