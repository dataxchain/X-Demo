package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *AssetboxChaincode) purchase(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	acct_id := args[0]
	assetbox_id := args[1]
	dt := args[2]

	// assetbox 정보 (소유자, 가격) 가져오기
	assetbox_bytes := t.getAssetbox(stub, []string{assetbox_id})
	var assetbox Assetbox
	err := json.Unmarshal(assetbox_bytes.Payload, &assetbox)

	// 구매자가 소유자에게 가격 만큼 코인 송금
	t.transferCoin(stub, []string{acct_id, assetbox.OwnerAcctId, assetbox.Price})

	// 거래 내역 저장
	objectType := "Purchasement"
	purchasementKey, _ := stub.CreateCompositeKey(objectType, []string{acct_id, assetbox_id})
	purchasement := Purchasement{acct_id, assetbox_id, dt}
	bytesAsPurchasement, _ := json.Marshal(purchasement)
	err = stub.PutState(purchasementKey, bytesAsPurchasement)
	if err != nil {
		shim.Error(err.Error())
	}

	//tracking 정보를 저장
	objectType = "Tracking"
	txId := stub.GetTxID()
	trackingKey, _ := stub.CreateCompositeKey(objectType, []string{txId})

	buyerAccount := t.getAccount(stub, []string{acct_id})

	tracking := Tracking{txId, dt, assetbox.Title, assetbox.Price,
		buyerAccount.Nickname, assetbox.OwnerId}
	bytesAsTracking, _ := json.Marshal(tracking)
	err = stub.PutState(trackingKey, bytesAsTracking)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) getAllTracking(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var trackings []Tracking
	objectType := "Tracking"

	tracikingResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer tracikingResultsIterator.Close()

	for tracikingResultsIterator.HasNext() {
		queryResponse, err := tracikingResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var tracking Tracking

		err = json.Unmarshal(queryResponse.Value, &tracking)
		if err != nil {
			return shim.Error(err.Error())
		}
		trackings = append(trackings, tracking)
	}
	ret, err := json.Marshal(trackings)
	return shim.Success(ret)
}

func (t *AssetboxChaincode) getAllPurchasement(
	stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var purchasements []Purchasement
	objectType := "Purchasement"
	acct_id := args[0]

	purchasementResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{acct_id})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer purchasementResultsIterator.Close()

	for purchasementResultsIterator.HasNext() {
		queryResponse, err := purchasementResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var purchasement Purchasement

		err = json.Unmarshal(queryResponse.Value, &purchasement)
		if err != nil {
			return shim.Error(err.Error())
		}
		purchasements = append(purchasements, purchasement)
	}

	ret, err := json.Marshal(purchasements)
	return shim.Success(ret)

}
