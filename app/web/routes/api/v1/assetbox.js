var express = require('express');
var router = express.Router();
const uuid = require('uuid');
const axios = require("axios");
const ipfs_api = require('ipfs-api');

var invokeCC = require('../../../../middleware/invoke_chaincode.js');
var queryCC = require('../../../../middleware/query_chaincode.js');


router.post('/', async function (req, res) {

    try {
        const {ownerId, ownerAcctId, title, description, price, images} = req.body;
        var chaincode_id = 'assetboxcc';
        var function_name = 'addAssetbox';
        var nickname = ownerId.split('@')[0];
        var time = new Date().toISOString().split('T')[0];
        var assetboxid = uuid.v4();
        var arg_list = [assetboxid, title, price, "0", nickname, ownerAcctId, "", time];
        console.log(arg_list);
        var assetids = [];
        var paths = [];

        invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list);

        const splitImages = images.split("***");

        for (var i = 0; i < splitImages.length; i++) {
            let base64Image = splitImages[i].split(';base64,')[1];
            let header = splitImages[i].split(';base64,')[0];
            let extname = header.split('/')[1];
            var assetid = uuid.v4();
            assetids.push(assetid);

            const ipfs = ipfs_api('ipfs.infura.io', '5001', {protocol: 'https'});

            content = ipfs.types.Buffer.from(new Buffer(base64Image, 'base64'));

            let file = await ipfs.files.add(content);
            //ipfs.files.cp('/ipfs/' + file[0].hash, '/' + title + ' ' + i + '.' + extname);
            //console.log(file[0].hash);
            const pullPath = "https://ipfs.io/ipfs/"+ file[0].hash;
            paths.push(pullPath);
        }

        var chaincode_id = 'assetboxcc';
        var function_name = 'addAsset';
        var arg_list = [assetids.join('##'), "", paths.join('##'), "", "", assetboxid];
        console.log(arg_list);
        invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list);

        res.json({
            code: 200,
            message: 'assetbox insert success',
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    }
});


router.post('/delete/:assetbox_id', async function(req, res){

    try {
        const { assetbox_id } = req.params;
        var chaincode_id = 'assetboxcc';
        var function_name = 'delAssetbox';
        var arg_list = [assetbox_id];

        await invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list);

        res.json({
            code: 200,
            message: 'assetbox delete success',
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    }
});

router.get('/uptest', function(req, res, next) {
    res.render('upload');
});

router.get('/', function (req, res, next) {
    var chaincode_id = 'assetboxcc';
    var function_name = 'getAllAssetbox';
    var arg_list = [];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            account: result
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });
});

router.get('/account/:acct_id', async function (req, res, next) {

    const {acct_id} = req.params;
    var chaincode_id = 'assetboxcc';
    var assetbox_list = [];
    var purchasement_list = [];

    await queryCC.query_chaincode(chaincode_id, "getAllAssetbox", []).then((result) => {
        if (result == null) {
            assetbox_list = [];
        } else {
            assetbox_list = result;
        }
        assetbox_list = result;
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });
    await queryCC.query_chaincode(chaincode_id, "getAllPurchasement", [acct_id]).then((result) => {
        if (result == null) {
            purchasement_list = [];
        } else {
            purchasement_list = result;
        }
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });
    // 각 assetbox에 대한 사용자의 구매 이력이 있는지 확인
    for (var idx = 0; idx < assetbox_list.length; idx++) {
        assetbox_list[idx].isPurchased = purchasement_list.some(
            purchasement => purchasement.assetboxId == assetbox_list[idx].id);
    }
    res.json({
        code: 200,
        account: assetbox_list
    });
});

router.post('/buy', async function (req, res, next) {

    const {acctid, assetboxid} = req.body;
    var chaincode_id = 'assetboxcc';
    var function_name = 'purchase';
    var dt = new Date().toISOString().slice(0, 10).replace(/-/g, '');
    var arg_list = [acctid, assetboxid, dt];

    await invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            message: 'buy success'
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });
});

router.get('/:assetboxid', function (req, res, next) {

    const {assetboxid} = req.params;
    var chaincode_id = 'assetboxcc';
    var function_name = 'getAssets';
    var arg_list = [assetboxid];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            account: result
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });

});

router.get('/tracking/:assetboxid', async function (req, res, next) {

    var chaincode_id = 'assetboxcc';
    const function_name = 'getAllTracking';
    var arg_list = [];

    try {
        const result = await queryCC.query_chaincode(chaincode_id, function_name, arg_list);
        console.log(result);
        for (let i = 0; i < result.length; i++) {
            let txId = result[i]["txId"];
            let tracking = result[i];
            const url = "http://35.168.16.248:8080/api/transaction/" +
            "247a1bcecf2c3844817c534fb7e2d7c281233da741aaa49c7d5c17d74d3b52b9/" + txId
            const response = await axios.get(url)
            const data = response.data;
            result[i]["transaction"] = data;

        }
        res.json({
            code: 200,
            account: result
        });
    } catch (err) {
        console.error(err);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    }

});
/*
    GET /assets/12345/meta/1234
 */
router.get('/assets/:assetid/meta/:acctid', function (req, res, next) {

    const {assetid, acctid} = req.params;
    console.log(assetid, acctid);
    var chaincode_id = 'assetboxcc';
    var function_name = 'getMetas';
    var arg_list = [assetid, acctid];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            account: result
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });

});

router.get('/assets/:assetid/user', function (req, res, next) {

    const {assetid} = req.params;
    console.log(assetid);
    var chaincode_id = 'assetboxcc';
    var function_name = 'getAllMetaAcctId';
    var arg_list = [assetid];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            account: result
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });

});

router.post('/assets/meta', function (req, res, next) {

    let {meta} = req.body;
    meta = JSON.parse(meta)
    let ids = [];
    let angle = [];
    let fill = [];
    let widths = [];
    let heights = [];
    let boxWidths = [];
    let boxHeights = [];
    let scaleFactors = [];
    let names = [];
    let left = [];
    let top = [];
    let assetId = meta[0].assetId;
    let ownerAcctId = meta[0].ownerAcctId;
    let ownerNickname = meta[0].ownerNickname;

    for (i = 0; i < meta.length; i++) {
        ids.push(uuid.v4());
        widths.push(meta[i].width);
        heights.push(meta[i].height);
        angle.push("0");
        fill.push("0");
        boxWidths.push(meta[i].box_width);
        boxHeights.push(meta[i].box_height);
        scaleFactors.push(meta[i].scaleFactor);
        names.push(meta[i].name);
        left.push(meta[i].left);
        top.push(meta[i].top);
    }

    console.log(ids, widths, heights, boxWidths, boxHeights, scaleFactors, names, assetId, ownerAcctId, ownerNickname);
    var chaincode_id = 'assetboxcc';
    var function_name = 'addMeta';
    var arg_list = [ids.join('##'), angle.join("##"), fill.join("##"), heights.join("##"), widths.join("##"),
        boxHeights.join("##"), boxWidths.join("##"), scaleFactors.join("##"), left.join("##"), top.join("##"), names.join("##"),
        assetId, ownerAcctId, ownerNickname]

    console.log(arg_list);

    invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        res.json({
            code: 200,
            message: 'meta insert success'
        });
    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });

});

module.exports = router;
