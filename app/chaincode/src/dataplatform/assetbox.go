package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
	"bytes"
)

type AssetboxChaincode struct {
	testMode bool
}

type Account struct {
	Acctid   string `json:"acctid"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Coin     string `json:"coin"`
}

type Meta struct {
	Id            string `json:"id"`
	Angle         string `json:"angle"`
	Fill          string `json:"fill"`
	Height        string `json:"height"`
	Width         string `json:"width"`
	MetaHeight    string `json:"metaHeight"`
	MetaWidth     string `json:"metaWidth"`
	ScaleFactor   string `json:"scaleFactor"`
	Left          string `json:"left"`
	Top           string `json:"top"`
	Label         string `json:"label"`
	AssetId       string `json:"assetId"`
	OwnerAcctId   string `json:"ownerAcctId"`
	OwnerNickname string `json:"ownerNickname"`
}

type Asset struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	ImgSrc string `json:"imgSrc"`
	Height string `json:"height"`
	Width  string `json:"width"`

	AssetboxId string `json:"assetboxId"`
}

type Purchasement struct {
	AcctId     string `json:"acctId"`
	AssetboxId string `json:"assetboxId"`
	Dt         string `json:"dt"`
}

type Assetbox struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Price        string `json:"price"`
	ProcessCount string `json:"processCount"`
	OwnerId      string `json:"ownerId"`
	OwnerAcctId  string `json:"ownerAcctId"`
	Thumbnail    string `json:"thumbnail"`
	Dt           string `json:"dt"`
}

type Tracking struct {
	TxId   string `json:"txId"`
	Dt     string `json:"dt"`
	Title  string `json:"title"`
	Price  string `json:"Price"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

func (t *AssetboxChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("assetbox Is Starting Up")
	//fmt.Println("init account state")

	//t.initState(stub, []string{"Account"})
	// Assetbox
	var assetboxes []Assetbox

	bytes, err := json.Marshal(assetboxes)

	if err != nil {
		return shim.Error("Error initializing assetbox.")
	}

	err = stub.PutState("assetbox", bytes)

	// Asset
	var assets []Asset

	bytes, err = json.Marshal(assets)

	if err != nil {
		return shim.Error("Error initializing assets.")
	}

	err = stub.PutState("asset", bytes)

	// meta
	var metas []Meta

	bytes, err = json.Marshal(metas)

	if err != nil {
		return shim.Error("Error initializing metas.")
	}

	err = stub.PutState("meta", bytes)

	// Purchasement
	var purchasements []Purchasement

	bytes, err = json.Marshal(purchasements)

	if err != nil {
		return shim.Error("Error initializing Purchasements.")
	}

	err = stub.PutState("purchasement", bytes)

	return shim.Success(nil)
}

func (t *AssetboxChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	if function == "addAccount" {
		return t.addAccount(stub, args)
	} else if function == "loginAccount" {
		return t.loginAccount(stub, args)
	} else if function == "getAllAssetbox" {
		return t.getAllAssetbox(stub, args)
	} else if function == "addAssetbox" {
		return t.addAssetbox(stub, args)
	} else if function == "getAssetbox" {
		return t.getAssetbox(stub, args)
	} else if function == "delAssetbox" {
		return t.delAssetbox(stub, args)
	} else if function == "addAsset" {
		return t.addAsset(stub, args)
	} else if function == "getAssets" {
		return t.getAssets(stub, args)
	} else if function == "addMeta" {
		return t.addMeta(stub, args)
	} else if function == "getMetas" {
		return t.getMetas(stub, args)
	} else if function == "getAllMetaAcctId" {
		return t.getAllMetaAcctId(stub, args)
	} else if function == "purchase" {
		return t.purchase(stub, args)
	} else if function == "getAllPurchasement" {
		return t.getAllPurchasement(stub, args)
	} else if function == "transferCoin" {
		return t.transferCoin(stub, args)
	} else if function == "isAccount" {
		return t.isAccount(stub, args)
	} else if function == "getCoin" {
		return t.getCoin(stub, args)
	} else if function == "getAllAccount" {
		return t.getAllAccount(stub, args)
	} else if function == "getAllTracking" {
		return t.getAllTracking(stub, args)
	}

	return shim.Error("Function with the name" + function + "does not exist.")
}

func (t *AssetboxChaincode) addAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	acctid := args[0]
	id := args[1]
	password := args[2]
	nickname := args[3]
	coinAmount := args[4]

	objectType := "Account"
	accountKey, _ := stub.CreateCompositeKey(objectType, []string{acctid})

	account := Account{acctid, id, password, nickname, coinAmount}
	bytes, _ := json.Marshal(account)
	err := stub.PutState(accountKey, bytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) addPointToAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	acctid := args[0]
	amount := args[1]

	objectType := "Account"
	accountKey, _ := stub.CreateCompositeKey(objectType, []string{acctid})

	bytes, err := stub.GetState(accountKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	var account Account
	json.Unmarshal(bytes, &account)
	currentAmount, _ := strconv.Atoi(account.Coin)
	addAmount, _ := strconv.Atoi(amount)

	account.Coin = strconv.Itoa(currentAmount + addAmount)

	ret, _ := json.Marshal(account)
	stub.PutState(accountKey, ret)

	return shim.Success(nil)
}

func (t *AssetboxChaincode) isAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	objectType := "Account"
	accountResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer accountResultsIterator.Close()

	isAccount := false
	var i int
	for i = 0; accountResultsIterator.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		queryResponse, err := accountResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var account Account
		err = json.Unmarshal(queryResponse.Value, &account)

		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to get %s: %s with error: %s", objectType, id, err))
		}
		if account.Id == id {
			isAccount = true
			break
		}
	}
	return shim.Success([]byte(strconv.FormatBool(isAccount)))

}

func (t *AssetboxChaincode) getAccount(stub shim.ChaincodeStubInterface, args []string) Account {
	var account Account
	acctid := args[0]
	objectType := "Account"
	accountKey, err := stub.CreateCompositeKey(objectType, []string{acctid})

	accountAsBytes, err := stub.GetState(accountKey)

	if err != nil {
		fmt.Printf("Unable to get account.")
	}

	err = json.Unmarshal(accountAsBytes, &account)
	return account
}

func (t *AssetboxChaincode) initState(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	objectType := args[0]

	accountResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer accountResultsIterator.Close()

	for accountResultsIterator.HasNext() {
		queryResponse, err := accountResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		stub.DelState(queryResponse.Key)

	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) getAllAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	objectType := "Account"

	accountResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer accountResultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for accountResultsIterator.HasNext() {
		queryResponse, err := accountResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllAccount:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (t *AssetboxChaincode) loginAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	id := args[0]
	password := args[1]

	objectType := "Account"
	var ret []byte
	accountResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer accountResultsIterator.Close()

	var i int
	for i = 0; accountResultsIterator.HasNext(); i++ {

		queryResponse, err := accountResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var account Account

		err = json.Unmarshal(queryResponse.Value, &account)

		if err != nil {
			return shim.Error(err.Error())
		}
		if account.Id == id && account.Password == password {
			account.Password = ""
			ret, _ = json.Marshal(account)
			break
		}

	}

	if ret == nil {
		jsonResp = "{\"Error\":\"Failed login : " + id + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(ret)
}

func (t *AssetboxChaincode) addAssetbox(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	title := args[1]
	price := args[2]
	processCount := args[3]
	ownerId := args[4]
	ownerAcctId := args[5]
	thumbnail := args[6]
	dt := args[7]

	objectType := "Assetbox"
	assetboxKey, _ := stub.CreateCompositeKey(objectType, []string{id})

	assetbox := Assetbox{id, title, price, processCount,
		ownerId, ownerAcctId, thumbnail, dt}
	bytes, _ := json.Marshal(assetbox)
	err := stub.PutState(assetboxKey, bytes)
	if err != nil {
		shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) getAssetbox(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]

	objectType := "Assetbox"
	assetboxKey, _ := stub.CreateCompositeKey(objectType, []string{id})

	bytesAsAssetbox, err := stub.GetState(assetboxKey)
	if err != nil {
		return shim.Error("Unable to get assetbox.")
	}

	return shim.Success(bytesAsAssetbox)
}

func (t *AssetboxChaincode) delAssetbox(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]

	objectType := "Assetbox"
	assetboxKey, _ := stub.CreateCompositeKey(objectType, []string{id})

	err := stub.DelState(assetboxKey)
	if err != nil {
		return shim.Error("Unable to delete assetbox.")
	}

	return shim.Success(nil)
}

func (t *AssetboxChaincode) getAllAssetbox(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	objectType := "Assetbox"
	var assetboxes []Assetbox
	assetboxResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer assetboxResultsIterator.Close()

	for assetboxResultsIterator.HasNext() {

		queryResponse, err := assetboxResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var assetbox Assetbox

		err = json.Unmarshal(queryResponse.Value, &assetbox)
		if err != nil {
			return shim.Error(err.Error())
		}

		//assetbox의 첫번째 asset의 이미지를 thunbnail로 저장함.
		objectType := "Asset"
		assetResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{assetbox.Id})
		if err != nil {
			return shim.Error(err.Error())
		}
		defer assetResultsIterator.Close()
		i := 0
		for i < 1 && assetResultsIterator.HasNext() {
			queryResponse, err := assetResultsIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			var asset Asset

			err = json.Unmarshal(queryResponse.Value, &asset)
			if err != nil {
				return shim.Error(err.Error())
			}
			assetbox.Thumbnail = asset.ImgSrc
			i++
		}
		assetboxes = append(assetboxes, assetbox)

	}

	ret, err := json.Marshal(assetboxes)

	return shim.Success(ret)
}

func (t *AssetboxChaincode) addMeta(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	bytes, err := stub.GetState("meta")

	if err != nil {
		return shim.Error("Unable to get asset.")
	}
	var metas []Meta
	err = json.Unmarshal(bytes, &metas)
	// Build JSON values
	id := strings.Split(args[0], "##")
	angle := strings.Split(args[1], "##")
	fill := strings.Split(args[2], "##")
	height := strings.Split(args[3], "##")
	width := strings.Split(args[4], "##")
	metaHeight := strings.Split(args[5], "##")
	metaWidth := strings.Split(args[6], "##")
	scaleFactor := strings.Split(args[7], "##")
	left := strings.Split(args[8], "##")
	top := strings.Split(args[9], "##")
	label := strings.Split(args[10], "##")
	assetId := args[11]
	ownerAcctId := args[12]
	ownerNickname := args[13]

	objectType := "Meta"

	//특정 사용자가 어떤 이미지에 기존에 넣은 메타정보를 지운다.
	metaResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{assetId, ownerAcctId})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer metaResultsIterator.Close()

	for metaResultsIterator.HasNext() {
		queryResponse, err := metaResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		stub.DelState(queryResponse.Key)
	}

	count := len(id)
	for i := 0; i < count; i++ {
		metaKey, _ := stub.CreateCompositeKey(objectType, []string{assetId, ownerAcctId, id[i]})
		meta := Meta{id[i], angle[i], fill[i], height[i],
			width[i], metaHeight[i], metaWidth[i],
			scaleFactor[i], left[i], top[i], label[i], assetId, ownerAcctId, ownerNickname}

		metaAsbytes, _ := json.Marshal(meta)
		err := stub.PutState(metaKey, metaAsbytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	}
	// 메타 등록한 사용자에게 100포인트 지급
	t.addPointToAccount(stub, []string{ownerAcctId, "100"})
	return shim.Success(nil)
}

func (t *AssetboxChaincode) getMetas(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var metas []Meta
	assetId := args[0]
	ownerAcctId := args[1]
	objectType := "Meta"

	metaResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{assetId, ownerAcctId})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer metaResultsIterator.Close()

	for metaResultsIterator.HasNext() {
		queryResponse, err := metaResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var meta Meta

		err = json.Unmarshal(queryResponse.Value, &meta)
		if err != nil {
			return shim.Error(err.Error())
		}
		metas = append(metas, meta)
	}

	ret, err := json.Marshal(metas)
	return shim.Success(ret)
}

func (t *AssetboxChaincode) getAllMetaAcctId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	metaUser := map[string]string{}
	assetId := args[0]
	objectType := "Meta"

	metaResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{assetId})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer metaResultsIterator.Close()

	for metaResultsIterator.HasNext() {
		queryResponse, err := metaResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var meta Meta

		err = json.Unmarshal(queryResponse.Value, &meta)
		if err != nil {
			return shim.Error(err.Error())
		}
		metaUser[meta.OwnerAcctId] = meta.OwnerNickname
	}
	ret, err := json.Marshal(metaUser)
	return shim.Success(ret)

}

func (t *AssetboxChaincode) addAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := strings.Split(args[0], "##")
	name := args[1]
	imgSrc := strings.Split(args[2], "##")
	height := args[3]
	width := args[4]
	assetboxId := args[5]

	objectType := "Asset"
	count := len(id)

	for i := 0; i < count; i++ {
		assetKey, _ := stub.CreateCompositeKey(objectType, []string{assetboxId, id[i]})
		asset := Asset{id[i], name, imgSrc[i], height,
			width, assetboxId}

		assetAsbytes, _ := json.Marshal(asset)
		err := stub.PutState(assetKey, assetAsbytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)

}

func (t *AssetboxChaincode) getAssets(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	assetboxid := args[0]
	objectType := "Asset"

	var assets []Asset
	assetResultsIterator, err := stub.GetStateByPartialCompositeKey(objectType, []string{assetboxid})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer assetResultsIterator.Close()

	for assetResultsIterator.HasNext() {
		queryResponse, err := assetResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var asset Asset

		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return shim.Error(err.Error())
		}
		assets = append(assets, asset)
	}

	ret, err := json.Marshal(assets)
	return shim.Success(ret)

}

func main() {
	err := shim.Start(new(AssetboxChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
