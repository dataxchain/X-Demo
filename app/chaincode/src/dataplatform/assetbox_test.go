package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"testing"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was", string(bytes), "and not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, function string, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(function), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	payload := string(res.Payload)
	if payload != value {
		fmt.Println("Query value", name, "was", payload, "and not", value, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvokeResult(t *testing.T, stub *shim.MockStub, args [][]byte, value string) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Invoke", args, "failed to get value")
		t.FailNow()
	}
	payload := string(res.Payload)
	if payload != value {
		fmt.Println("Invoke value", args, "was", payload, "and not", value, "as expected")
		t.FailNow()
	}
}

func getInitArguments() [][]byte {
	return [][]byte{[]byte("init"),
		[]byte("fisrt"),
		[]byte("second")}
}

/*
func TestAssetbox(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAsset'
	id := "0"
	account := "1"
	name := "hello.txt"
	fileHash := "hashabcd"
	encryptionKey := "key123"

	checkInvoke(t, stub, [][]byte{[]byte("addAsset"), []byte(id),
		[]byte(account), []byte(name), []byte(fileHash), []byte(encryptionKey)})

	asset := &Asset{
		id, account, name, fileHash, encryptionKey}
	assets := []*Asset{asset}
	assetsBytes, _ := json.Marshal(assets)
	assetKey := "assetbox_assets"
	checkState(t, stub, assetKey, string(assetsBytes))

	expectedResp := string(assetsBytes)
	checkQuery(t, stub, "browseAsset", account, expectedResp)
}

func TestAssetboxGetAsset(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAsset'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAsset"),
			[]byte("0"),         // id
			[]byte("1"),         // account
			[]byte("hello.txt"), // name
			[]byte("hashabcd"),  // fileHash
			[]byte("key123")}) // encryptionKey

	checkInvoke(t, stub,
		[][]byte{[]byte("addAsset"),
			[]byte("1"),        // id
			[]byte("2"),        // account
			[]byte("hihi.txt"), // name
			[]byte("hashzzzz"), // fileHash
			[]byte("key456")}) // encryptionKey

	asset := &Asset{
		"1", "2", "hihi.txt", "hashzzzz", "key456"}
	assetBytes, _ := json.Marshal(asset)

	expectedResp := string(assetBytes)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAsset"),
			[]byte("1"), // id
			[]byte("2")}, // account_id
		expectedResp)
}
*/
func TestLoginAccount2(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())


	for i := 0 ; i < 10 ; i++ {
		checkInvoke(t, stub,
			[][]byte{[]byte("addAccount"),
				[]byte(strconv.Itoa(i)),        // acctid
				[]byte("test" + strconv.Itoa(i)),  // id
				[]byte("pass1234"), // password
				[]byte("moons"),    // nickname
				[]byte("1000")}) // coin

		account := &Account{
			Acctid: strconv.Itoa(i), Id: "test" + strconv.Itoa(i), Password: "", Nickname: "moons", Coin: "1000"}
		successResult, _ := json.Marshal(account)
		successExpectedResp := string(successResult)

		checkInvokeResult(t, stub,
			[][]byte{[]byte("loginAccount"),
				[]byte("test" + strconv.Itoa(i)), // id
				[]byte("pass1234")}, // password
			successExpectedResp)
	}

}

func TestLoginAccount(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	checkInvokeResult(t, stub,
		[][]byte{[]byte("isAccount"),
			[]byte("moonsub")}, // id
		"false")

	// Invoke 'addAsset'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("1"),        // acctid
			[]byte("moonsub"),  // id
			[]byte("pass1234"), // password
			[]byte("moons"),    // nickname
			[]byte("1000")}) // coin

	checkInvokeResult(t, stub,
		[][]byte{[]byte("isAccount"),
			[]byte("moonsub")}, // id
		"true")

	checkInvokeResult(t, stub,
		[][]byte{[]byte("isAccount"),
			[]byte("idid")}, // id
		"false")

	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("2"),        // acctid
			[]byte("idid"),     // id
			[]byte("pass1234"), // password
			[]byte("idid"),     // nickname
			[]byte("2000")}) // coin

	checkInvokeResult(t, stub,
		[][]byte{[]byte("isAccount"),
			[]byte("idid")}, // id
		"true")

	checkInvokeResult(t, stub,
		[][]byte{[]byte("isAccount"),
			[]byte("idid")}, // id
		"true")

	account := &Account{
		Acctid: "1", Id: "moonsub", Password: "", Nickname: "moons", Coin: "1000"}
	successResult, _ := json.Marshal(account)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("loginAccount"),
			[]byte("moonsub"), // id
			[]byte("pass1234")}, // password
		successExpectedResp)

	account2 := &Account{
		Acctid: "2", Id: "idid", Password: "", Nickname: "idid", Coin: "2000"}
	successResult2, _ := json.Marshal(account2)

	successExpectedResp2 := string(successResult2)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("loginAccount"),
			[]byte("idid"), // id
			[]byte("pass1234")}, // password
		successExpectedResp2)

}

func TestGetDelAssetbox(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAssetbox'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAssetbox"),
			[]byte("1234-1231"), // id
			[]byte("title1"),    // title
			[]byte("50"),        //price
			[]byte("0"),         //processCount
			[]byte("moonsub"),   //ownerId
			[]byte("1"),         //ownerAcctId
			[]byte(""),          //thumbnail
			[]byte("2018-10-01")}) //dt

	checkInvoke(t, stub,
		[][]byte{[]byte("addAssetbox"),
			[]byte("1234-5232"), // id
			[]byte("title2"),    // title
			[]byte("120.7"),     //price
			[]byte("3"),         //processCount
			[]byte("d1d1"),      //ownerId
			[]byte("2"),         //ownerAcctId
			[]byte(""),          //thumbnail
			[]byte("2018-10-02")}) //dt

	assetbox1 := Assetbox{
		Id: "1234-1231", Title: "title1", Price: "50", ProcessCount: "0",
		OwnerId: "moonsub", OwnerAcctId: "1", Dt: "2018-10-01"}

	assetbox2 := Assetbox{
		Id: "1234-5232", Title: "title2", Price: "120.7", ProcessCount: "3",
		OwnerId: "d1d1", OwnerAcctId: "2", Dt: "2018-10-02"}

	result1, _ := json.Marshal(assetbox1)
	checkInvokeResult(t, stub, [][]byte{
		[]byte("getAssetbox"), []byte("1234-1231")}, string(result1))

	result2, _ := json.Marshal(assetbox2)
	checkInvokeResult(t, stub, [][]byte{
		[]byte("getAssetbox"), []byte("1234-5232")}, string(result2))

	// test delete assetbox
	checkInvoke(t, stub, [][]byte{
		[]byte("delAssetbox"), []byte("1234-5232")})

	var assetboxes []Assetbox
	assetboxes = append(assetboxes, assetbox1)
	successResult, _ := json.Marshal(assetboxes)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAllAssetbox")}, successExpectedResp)
}

func TestGetAllAssetbox(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAssetbox'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAssetbox"),
			[]byte("1234-1231"), // id
			[]byte("title1"),    // title
			[]byte("50"),        //price
			[]byte("0"),         //processCount
			[]byte("moonsub"),   //ownerId
			[]byte("1"),         //ownerAcctId
			[]byte(""),          //thumbnail
			[]byte("2018-10-01")}) //dt

	checkInvoke(t, stub,
		[][]byte{[]byte("addAsset"),
			[]byte("123-123##234-234##456-456"), //id
			[]byte(""),                          //name
			[]byte("/static/1_image1##/static/1_image2##/static/1_image3"), //imgSrc
			[]byte(""), //height
			[]byte(""), //width
			[]byte("1234-1231")}) //assetboxId

	var assetboxes []Assetbox
	assetbox1 := Assetbox{Id: "1234-1231", Title: "title1", Price: "50", ProcessCount: "0",
		OwnerId: "moonsub", OwnerAcctId: "1", Thumbnail: "/static/1_image1", Dt: "2018-10-01"}
	assetboxes = append(assetboxes, assetbox1)
	successResult, _ := json.Marshal(assetboxes)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAllAssetbox")}, // password
		successExpectedResp)

}

func TestGetAssets(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAsset'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAsset"),
			[]byte("123-123##234-234##456-456"), //id
			[]byte(""),                          //name
			[]byte("/static/1_image1##/static/1_image2##/static/1_image3"), //imgSrc
			[]byte(""), //height
			[]byte(""), //width
			[]byte("1")}) //assetboxId

	var assets []Asset
	asset1 := Asset{Id: "123-123", Name: "", ImgSrc: "/static/1_image1", Height: "", Width: "", AssetboxId: "1"}
	asset2 := Asset{Id: "234-234", Name: "", ImgSrc: "/static/1_image2", Height: "", Width: "", AssetboxId: "1"}
	asset3 := Asset{Id: "456-456", Name: "", ImgSrc: "/static/1_image3", Height: "", Width: "", AssetboxId: "1"}
	assets = append(assets, asset1)
	assets = append(assets, asset2)
	assets = append(assets, asset3)
	successResult, _ := json.Marshal(assets)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAssets"),
			[]byte("1")}, // assetboxId
		successExpectedResp)

	var emptyAsset []Asset

	emptyResult, _ := json.Marshal(emptyAsset)

	emptyExpectedResp := string(emptyResult)
	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAssets"),
			[]byte("10")}, // assetboxId
		emptyExpectedResp)

}

func TestGetMetas(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAccount'
	checkInvoke(t, stub,
		[][]byte{[]byte("addAccount"),
			[]byte("2"),        // acctid
			[]byte("moonsub"),  // id
			[]byte("pass1234"), // password
			[]byte("moons"),    // nickname
			[]byte("1000")}) // coin

	// Invoke 'addAsset'
	checkInvoke(t, stub,
		[][]byte{[]byte("addMeta"),
			[]byte("123-123##234-234##456-456"),                               //id
			[]byte("10##10##10"),                                              //angle
			[]byte("rgba(255,0,0,0,0)##rgba(255,0,0,0,0)##rgba(255,0,0,0,0)"), //Fill
			[]byte("20##20##20"),                                              //Height
			[]byte("30##30##30"),                                              //Width
			[]byte("40##40##40"),                                              //MetaHeight
			[]byte("50##50##50"),                                              //MetaWidth
			[]byte("60##60##60"),                                              //ScaleFactor
			[]byte("70##70##70"),                                              //Left
			[]byte("80##80##80"),                                              //Top
			[]byte("90##90##90"),                                              //Label
			[]byte("1"),                                                       //AssetId
			[]byte("2"),                                                       //OwnerAcctId
			[]byte("m.kwon")}) //OwnerNickname})

	var metas []Meta
	meta1 := Meta{Id: "123-123", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	meta2 := Meta{Id: "234-234", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	meta3 := Meta{Id: "456-456", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	metas = append(metas, meta1)
	metas = append(metas, meta2)
	metas = append(metas, meta3)
	successResult, _ := json.Marshal(metas)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getMetas"),
			[]byte("1"), //assetId
			[]byte("2")}, // OwnerAcctId
		successExpectedResp)

	checkInvoke(t, stub,
		[][]byte{[]byte("addMeta"),
			[]byte("123-123##234-234##456-456"),                               //id
			[]byte("10##10##10"),                                              //angle
			[]byte("rgba(255,0,0,0,0)##rgba(255,0,0,0,0)##rgba(255,0,0,0,0)"), //Fill
			[]byte("20##20##20"),                                              //Height
			[]byte("30##30##30"),                                              //Width
			[]byte("40##40##40"),                                              //MetaHeight
			[]byte("50##50##50"),                                              //MetaWidth
			[]byte("60##60##60"),                                              //ScaleFactor
			[]byte("70##70##70"),                                              //Left
			[]byte("80##80##80"),                                              //Top
			[]byte("90##90##90"),                                              //Label
			[]byte("1"),                                                       //AssetId
			[]byte("2"),                                                       //OwnerAcctId
			[]byte("m.kwon")}) //OwnerNickname

	var new_metas []Meta
	new_metas1 := Meta{Id: "123-123", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	new_metas2 := Meta{Id: "234-234", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	new_metas3 := Meta{Id: "456-456", Angle: "10", Fill: "rgba(255,0,0,0,0)", Height: "20", Width: "30",
		MetaHeight: "40", MetaWidth: "50", ScaleFactor: "60", Left: "70", Top: "80", Label: "90", AssetId: "1", OwnerAcctId: "2", OwnerNickname: "m.kwon"}
	new_metas = append(new_metas, new_metas1)
	new_metas = append(new_metas, new_metas2)
	new_metas = append(new_metas, new_metas3)
	new_successResult, _ := json.Marshal(new_metas)

	newSuccessExpectedResp := string(new_successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getMetas"),
			[]byte("1"), //assetId
			[]byte("2")}, // OwnerAcctId
		newSuccessExpectedResp)

	checkInvokeResult(t, stub,
		[][]byte{
			[]byte("getCoin"),
			[]byte("2")},
		string("1200"))
}

func TestGetAllMetaAcctId(t *testing.T) {
	scc := new(AssetboxChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'addAsset'
	checkInvoke(t, stub,
		[][]byte{[]byte("addMeta"),
			[]byte("123-123##234-234##456-456"),                               //id
			[]byte("10##10##10"),                                              //angle
			[]byte("rgba(255,0,0,0,0)##rgba(255,0,0,0,0)##rgba(255,0,0,0,0)"), //Fill
			[]byte("20##20##20"),                                              //Height
			[]byte("30##30##30"),                                              //Width
			[]byte("40##40##40"),                                              //MetaHeight
			[]byte("50##50##50"),                                              //MetaWidth
			[]byte("60##60##60"),                                              //ScaleFactor
			[]byte("70##70##70"),                                              //Left
			[]byte("80##80##80"),                                              //Top
			[]byte("90##90##90"),                                              //Label
			[]byte("1"),                                                       //AssetId
			[]byte("2"),                                                       //OwnerAcctId
			[]byte("m.kwon")}) //OwnerNickname

	checkInvoke(t, stub,
		[][]byte{[]byte("addMeta"),
			[]byte("123-123##234-234##456-456"),                               //id
			[]byte("10##10##10"),                                              //angle
			[]byte("rgba(255,0,0,0,0)##rgba(255,0,0,0,0)##rgba(255,0,0,0,0)"), //Fill
			[]byte("20##20##20"),                                              //Height
			[]byte("30##30##30"),                                              //Width
			[]byte("40##40##40"),                                              //MetaHeight
			[]byte("50##50##50"),                                              //MetaWidth
			[]byte("60##60##60"),                                              //ScaleFactor
			[]byte("70##70##70"),                                              //Left
			[]byte("80##80##80"),                                              //Top
			[]byte("90##90##90"),                                              //Label
			[]byte("1"),                                                       //AssetId
			[]byte("5"),                                                       //OwnerAcctId
			[]byte("m.test")}) //OwnerNickname

	acctIds := map[string]string{}
	acctIds["2"] = "m.kwon"
	acctIds["5"] = "m.test"
	successResult, _ := json.Marshal(acctIds)

	successExpectedResp := string(successResult)

	checkInvokeResult(t, stub,
		[][]byte{[]byte("getAllMetaAcctId"),
			[]byte("1")}, //assetId
		successExpectedResp)

}
