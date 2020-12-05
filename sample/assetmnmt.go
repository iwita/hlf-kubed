package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type AssetMgr struct {
}

//define organization asset information, the record can be trace in bloackchain
type OrgAsset struct {
	Id        string `json:”id”`        //the assetId
	AssetType string `json:”assetType”` //type of device
	Status    string `json:”status”`    //status of asset
	Location  string `json:”location”`  //device location
	DeviceId  string `json:”deviceId”`  //DeviceId
	Comment   string `json:”comment”`   //comment
	From      string `json:”from”`      //from
	To        string `json:”to”`        //to
}

// func (c *AssetMgr) Init(stub shim.ChaincodeStubInterface) pb.Response {

// 	return shim.Success(nil)

// }

func (c *AssetMgr) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetStringArgs()
	if len(args) != 3 {

		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	assetId := args[0]
	assetType := args[1]
	deviceId := args[2]

	//create asset
	assetData := OrgAsset{
		Id:        assetId,
		AssetType: assetType,
		Status:    "START",
		Location:  "N/A",
		DeviceId:  deviceId,
		Comment:   "Initialized asset",
		From:      "N/A",
		To:        "N/A",
	}

	assetBytes, _ := json.Marshal(assetData)
	assetErr := stub.PutState(assetId, assetBytes)
	if assetErr != nil {

		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)
}

func (c *AssetMgr) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "Order" {
		return c.Order(stub, args)
	} else if function == "Ship" {
		return c.Ship(stub, args)
	} else if function == "Distribute" {
		return c.Distribute(stub, args)
	} else if function == "query" {
		return c.query(stub, args)
	} else if function == "getHistory" {
		return c.getHistory(stub, args)
	}
	return shim.Error("Invalid function name")
}

func (c *AssetMgr) Order(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return c.UpdateAsset(stub, args, "ORDER", "SCHOOL", "OEM")
}

func (c *AssetMgr) Ship(stub shim.ChaincodeStubInterface, args []string) pb.Response {
}

func (c *AssetMgr) Distribute(stub shim.ChaincodeStubInterface, args []string) pb.Response {

}

func (c *AssetMgr) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	type AuditHistory struct {
		TxId  string   `json:”txId”`
		Value OrgAsset `json:”value”`
	}
	var history []AuditHistory
	var orgAsset OrgAsset
	assetId := args[0]

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(assetId)
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		var tx AuditHistory
		tx.TxId = historyData.TxId
		json.Unmarshal(historyData.Value, &orgAsset)
		tx.Value = orgAsset           //copy orgAsset over
		history = append(history, tx) //add this tx to the list

	}
}

func (c *AssetMgr) UpdateAsset(stub shim.ChaincodeStubInterface, args []string, currentStatus string, from string, to string) pb.Response {
	assetId := args[0]
	comment := args[1]
	location := args[2]
	assetBytes, err := stub.GetState(assetId)
	orgAsset := OrgAsset{}
	if currentStatus == "ORDER" && orgAsset.Status != "START" {
		return shim.Error(err.Error())

	} else if currentStatus == "SHIP" && orgAsset.Status != "ORDER" {

	} else if currentStatus == "DISTRIBUTE" && orgAsset.Status != "SHIP" {

	}
	orgAsset.Comment = comment
	orgAsset.Status = currentStatus
	orgAsset0, _ := json.Marshal(orgAsset)
	err = stub.PutState(assetId, orgAsset0)
	return shim.Success(orgAsset0)

}

func main() {

	err := shim.Start(new(AssetMgr))
	if err != nil {
		fmt.Printf("Error creating new AssetMgr Contract: %s", err)
	}

}
