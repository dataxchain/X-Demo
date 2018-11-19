package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func (t *AssetboxChaincode) transferCoin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	objectType := "Account"
	// 보내는 사용자와 받는 사용자 가져오기
	get_account := func(acct_id string) Account {
		var selected_account Account
		accountKey, _ := stub.CreateCompositeKey(objectType, []string{acct_id})
		bytes, _ := stub.GetState(accountKey)
		json.Unmarshal(bytes, &selected_account)
		return selected_account
	}

	from_acct_id := args[0]
	to_acct_id := args[1]

	from_account := get_account(from_acct_id)
	to_account := get_account(to_acct_id)
	coin_amount, _ := strconv.Atoi(args[2])

	from_coin, _ := strconv.Atoi(from_account.Coin)
	to_coin, _ := strconv.Atoi(to_account.Coin)

	// 계산
	from_account.Coin = strconv.Itoa(from_coin - coin_amount)
	to_account.Coin = strconv.Itoa(to_coin + coin_amount)

	// 보내는 사용자와 받는 사용자는 변경된 내용을 적용하고 그렇지 않은 경우는 그대로 둠
	fromAccountKey, _ := stub.CreateCompositeKey(objectType, []string{from_acct_id})
	fromAsBytes, _ := json.Marshal(from_account)
	err := stub.PutState(fromAccountKey, fromAsBytes)
	if err != nil{
		shim.Error("put state errored")
	}
	toAccountKey, _ := stub.CreateCompositeKey(objectType, []string{to_acct_id})
	toAsBytes, _ := json.Marshal(to_account)
	err = stub.PutState(toAccountKey, toAsBytes)
	if err != nil{
		shim.Error("put state errored")
	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) getCoin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	acctid := args[0]
	objectType := "Account"
	accountKey, err := stub.CreateCompositeKey(objectType, []string{acctid})

	accountAsBytes, err := stub.GetState(accountKey)
	var account Account
	err = json.Unmarshal(accountAsBytes, &account)

	if err != nil {
		return shim.Error("Unable to get account.")
	}

	return shim.Success([]byte(account.Coin))
}