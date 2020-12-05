package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"

	"github.com/stretchr/testify/assert"
)

func TestInvokeInitSmartContract(t *testing.T) {
	fmt.Println("Entering TestInvokeInitSmartContract")

	// Instantiate mockStub
	stub := shimtest.NewMockStub("mockStub", new(SmartContract))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	var expected = "hello"

	// The first parameter is a uuid
	result := stub.MockInvoke("001", [][]byte{[]byte("helloF")})

	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState("res")
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	assert.Equal(t, string(valAsbytes), expected)
}
