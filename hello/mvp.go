package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SmartContract struct {
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "helloF" {
		return t.hello(stub, args)
	}
	return shim.Error("Invalid invoke function name.")
}

func (t *SmartContract) hello(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if err := stub.PutState("res", []byte("hello")); err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
