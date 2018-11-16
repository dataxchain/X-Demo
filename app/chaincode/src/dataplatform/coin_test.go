package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestTransferCoin(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// 첫 번째 사용자에게 80 두 번재 사용자에게 210 코인을 준 뒤 30만큼 송금하여 결과 확인

	// Init
	checkInit(t, stub, getInitArguments())

	checkInvoke(t, stub, [][]byte{
		[]byte("addAccount"), []byte("1"), []byte("first"), []byte("aaa"),
		[]byte("one"), []byte("80")})

	checkInvoke(t, stub, [][]byte{
		[]byte("addAccount"), []byte("2"), []byte("second"), []byte("bbb"),
		[]byte("two"), []byte("210")})

	test_account := func(acct_id string, user_id string, password string,
		nickname string, coin_amount string) {

		checkInvokeResult(t, stub, [][]byte{
			[]byte("getCoin"), []byte(acct_id)}, string(coin_amount))
	}

	test_account("1", "first",
		"aaa", "one", "80")

	test_account("2", "second",
		"bbb", "two", "210")

	checkInvoke(t, stub, [][]byte{
		[]byte("transferCoin"), []byte("2"), []byte("1"), []byte("30")})

	test_account("1", "first",
		"aaa", "one", "110")

	test_account("2", "second",
		"bbb", "two", "180")
}
