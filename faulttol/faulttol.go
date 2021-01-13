package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ServerConfig struct {
	CCID    string
	Address string
}

type SmartContract struct {
	contractapi.Contract
}

type Resources struct {
	NetBw  float32
	CPU    float32
	Memory float32
}

type Device struct {
	Kind     string
	Uuid     string
	Org      string
	Capacity *Resources
	Used     *Resources
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Device
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	devices := []Device{
		{
			Kind: "Edge",
			Uuid: "123",
			Org:  "org1",
			Capacity: &Resources{
				CPU: 100, Memory: 1000,
			},
			Used: &Resources{
				CPU: 0, Memory: 0,
			},
		},
		{
			Kind: "Edge",
			Uuid: "456",
			Org:  "org2",
			Capacity: &Resources{
				CPU: 100, Memory: 1000,
			},
			Used: &Resources{
				CPU: 0, Memory: 0,
			},
		},
	}

	for i, dev := range devices {
		devAsBytes, _ := json.Marshal(dev)
		if err := ctx.GetStub().PutState("Device"+strconv.Itoa(i), devAsBytes); err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}

	}
	return nil
}

func (s *SmartContract) QueryAllDevices(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		dev := new(Device)
		_ = json.Unmarshal(queryResponse.Value, dev)

		queryResult := QueryResult{Key: queryResponse.Key, Record: dev}
		results = append(results, queryResult)
	}
	return results, nil
}

func (s *SmartContract) QueryDev(ctx contractapi.TransactionContextInterface, id string) (*Device, error) {
	carAsBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	dev := new(Device)
	_ = json.Unmarshal(carAsBytes, dev)

	return dev, nil
}

func (s *SmartContract) UpdateDeviceUsedResources(ctx contractapi.TransactionContextInterface, id string, cpu, mem, bw float32) error {
	dev, err := s.QueryDev(ctx, id)

	if err != nil {
		return err
	}
	dev.Used.CPU = cpu
	dev.Used.Memory = mem
	dev.Used.NetBw = bw
	devAsBytes, _ := json.Marshal(dev)

	return ctx.GetStub().PutState(id, devAsBytes)
}

func (s *SmartContract) GetDeviceAvailableResources(ctx contractapi.TransactionContextInterface, uuid string) (*Resources, error) {
	devAsBytes, err := ctx.GetStub().GetState(uuid)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if devAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", uuid)
	}

	dev := new(Device)
	_ = json.Unmarshal(devAsBytes, dev)
	available := &Resources{}
	available.CPU = dev.Capacity.CPU - dev.Used.CPU
	available.NetBw = dev.Capacity.NetBw - dev.Used.NetBw
	available.Memory = dev.Capacity.Memory - dev.Used.Memory
	return available, nil
}

func (s *SmartContract) GetAllDevicesAvailableResources(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	// TO IMPLEMENT
	return []QueryResult{}, nil
}

func main() {
	// See chaincode.env.example
	config := ServerConfig{
		CCID:    os.Getenv("CHAINCODE_CCID"),
		Address: os.Getenv("CHAINCODE_ADDRESS"),
	}

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create faulttol chaincode: %s", err.Error())
		return
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting faulttol chaincode: %s", err.Error())
	}
}
