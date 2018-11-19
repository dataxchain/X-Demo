package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestPurchaseAssetbox(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("1"),        // acctid
			[]byte("test123"),  // id
			[]byte("pass1234"), // password
			[]byte("test1"),    // nickname
			[]byte("1000")}) // coin

	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("2"),        // acctid
			[]byte("test456"),  // id
			[]byte("pass1234"), // password
			[]byte("test2"),    // nickname
			[]byte("1000")}) // coin

	checkInvoke(t, stub,
		[][]byte{[]byte("addAssetbox"),
			[]byte("1"),      // id
			[]byte("title1"), // title
			[]byte("100"),    //price
			[]byte("0"),      //processCount
			[]byte("test1"),  //ownerId
			[]byte("1"),      //ownerAcctId
			[]byte(""),       //thumbnail
			[]byte("2018-10-01")}) //dt

	// 구매
	purchase := func(acct_id string, assetbox_id string, dt string) {
		checkInvoke(t, stub,
			[][]byte{
				[]byte("purchase"),
				[]byte(acct_id),
				[]byte(assetbox_id),
				[]byte(dt)})
	}
	purchase("2", "1", "20180102")


	// 첫 번째 사용자가 구매한 assetbox 목록 확인

	var purchasements []Purchasement

	purchasements = append(purchasements, Purchasement{AcctId: "2", AssetboxId: "1", Dt: "20180102"})

	successResult, _ := json.Marshal(purchasements)
	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getCoin"),
			[]byte("1")},
		string("1100"))

	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getCoin"),
			[]byte("2")},
		string("900"))

	var trackings []Tracking

	tracking := Tracking {"1","20180102","title1","100","test2","test1"}
	trackings = append(trackings, tracking)

	successResult, _ = json.Marshal(trackings)
	successExpectedResp = string(successResult)
	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getAllTracking")},
		successExpectedResp)

}

func TestPurchase(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("1"),        // acctid
			[]byte("test123"),  // id
			[]byte("pass1234"), // password
			[]byte("test1"),    // nickname
			[]byte("1000")}) // coin

	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("2"),        // acctid
			[]byte("test456"),  // id
			[]byte("pass1234"), // password
			[]byte("test2"),    // nickname
			[]byte("1000")}) // coin

	// 구매
	purchase := func(acct_id string, assetbox_id string, dt string) {
		checkInvoke(t, stub,
			[][]byte{
				[]byte("purchase"),
				[]byte(acct_id),
				[]byte(assetbox_id),
				[]byte(dt)})
	}
	purchase("2", "4", "20180102")
	purchase("1", "4", "20180102")
	purchase("1", "5", "20180102")
	purchase("1", "6", "20180102")
	purchase("2", "7", "20180102")

	// 첫 번째 사용자가 구매한 assetbox 목록 확인

	var purchasements []Purchasement

	purchasements = append(purchasements, Purchasement{AcctId: "1", AssetboxId: "4", Dt: "20180102"})
	purchasements = append(purchasements, Purchasement{AcctId: "1", AssetboxId: "5", Dt: "20180102"})
	purchasements = append(purchasements, Purchasement{AcctId: "1", AssetboxId: "6", Dt: "20180102"})

	successResult, _ := json.Marshal(purchasements)
	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getAllPurchasement"),
			[]byte("1")},
		successExpectedResp)

	var purchasements_2 []Purchasement

	purchasements_2 = append(purchasements_2, Purchasement{AcctId: "2", AssetboxId: "4", Dt: "20180102"})
	purchasements_2 = append(purchasements_2, Purchasement{AcctId: "2", AssetboxId: "7", Dt: "20180102"})

	successResult_2, _ := json.Marshal(purchasements_2)
	successExpectedResp_2 := string(successResult_2)

	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getAllPurchasement"),
			[]byte("2")},
		successExpectedResp_2)


}
