
package main

import (
	"errors"
	//"encoding/json"	
	//"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


type account struct{
RollNumber int `json:"rollnumber"`
Name string `json:"name"`
Percent int `json:"percent"`
Year int `json:"year"`
College string `json:"college"`
}



// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("RollNumber", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "addRecord" {
		return t.addRecord(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) addRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte,error){
		
	var err error
	var rollnumber string =args[0]	
	if len(args)!=5 {
		return nil, errors.New("Incorrect number of args, expected 5 for record entry")
	}
	

	str:=`{"rollnumber": `+args[0]+`, "name": "`+args[1]+`", "percent": `+args[2]+`, "year":`+args[3]+`, "college":"`+args[4]+`"}`
	err=stub.PutState(rollnumber,[]byte(str))
	if err!=nil {
		return nil, err
	}
	 return nil,nil
	
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface,function string,args []string) ([]byte, error){
	

	if function=="getInfo" {
		return t.getInfo(stub,args)
	}
	
	fmt.Println("didnt find any function"+function)
	
	return nil,errors.New("unknown query")
}


func (t *SimpleChaincode) getInfo(stub shim.ChaincodeStubInterface,args []string)([]byte, error){
	
	var err error
	if len(args)!=1 {
		return nil, errors.New("wrong number of arguments to get info")
	}	

	valAsbytes,err := stub.GetState(args[0])
	if err!=nil {
		return nil, errors.New("couldnt get the record, check id sent")
	}	
	
	return valAsbytes, nil
	
}




