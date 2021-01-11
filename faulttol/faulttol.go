package faulttol

import (
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
	Capacity *Resources
	Used     *Resources
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Resources
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// TO IMPLEMENT
	return nil
}

func (s *SmartContract) UpdateDeviceUsedResources(ctx contractapi.TransactionContextInterface, Uuid string, rss *Resources) error {
	// TO IMPLEMENT
	return nil
}

func (s *SmartContract) GetDeviceAvailableResources(ctx contractapi.TransactionContextInterface, Uuid string) (*Resources, error) {
	// TO IMPLEMENT
	return nil, nil
}

func (s *SmartContract) GetAllDevicesAvailableResources(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	// TO IMPLEMENT
	return []QueryResult{}, nil
}

func main() {

}
