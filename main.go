package main

import (
	"fmt"
	"hlf-kubed/faulttol"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(new(faulttol.SmartContract))
	if err != nil {
		fmt.Printf("Error creating faulttol chaincode: %v\n", err)
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting faulttol chaincode: %v", err)
	}

}
