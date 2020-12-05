# Sample Chaincode in Go

## Some info first:
Source: https://developer.ibm.com/recipes/tutorials/writing-hyperledger-fabric-chaincode-using-go-programming-language/

- Chaincode needs to be installed on each endorsing peer node that runs on a secured Docker
- Chaincode in Hyperledger Fabric is similar to smart contracts
- The application can interact with the blockchain by invoking chaincode to manage the ledger state and keep the transaction record in the ledger.
- Every chaincode program must implement the Chaincode interface

```golang
type Chaincode interface {

    // create an initial state and the data initialization after the chaincode container has been established for the first time
    Init(stub ChaincodeStubInterface) pb.Response 

    //The Invoke method is called to interact with the ledger (to query or update the asset) in the proposed transaction
    Invoke(stub ChaincodeStubInterface) pb.Response
}
```

ChaincodeStubInterface provides the API for apps to access and modify their ledgers

```golang
type ChaincodeStubInterface interface {

InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response

GetState(key string) ([]byte, error)

PutState(key string, value []byte) error

DelState(key string) error

GetQueryResult(query string) (StateQueryIteratorInterface, error)

GetTxTimestamp() (*timestamp.Timestamp, error)

GetTxID() string

GetChannelID() string

}
```

