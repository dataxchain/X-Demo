var express = require('express');
const uuid = require('uuid');

var router = express.Router();

var queryCC = require('../../../../middleware/query_chaincode.js');
var invokeCC = require('../../../../middleware/invoke_chaincode.js');

router.post('/', function (req, res, next) {

    const {account, password} = req.body;
    const nickname = account.split("@")[0];
    var chaincode_id = 'assetboxcc';
    var function_name = 'isAccount';
    var arg_list = [account]

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        console.log(result);
        if (result != true) {
            var function_name = 'addAccount';
            const acct_id = uuid.v4()
            var arg_list = [acct_id, account, password, nickname, "10000"];
            console.log(arg_list);

            invokeCC.invoke_chaincode(chaincode_id, function_name, arg_list).then((result) => {
                res.json({
                    code   : 200,
                    message: 'account join success'
                });
            }, (err) => {
                console.error(err);
                res.status(500).json({
                    code   : 500,
                    message: 'server error',
                });
            });
        } else {
            res.json({
                code   : 401,
                message: 'id already exists'
            });
        }
    });

});

router.get('/coin/:acctid', function (req, res, next) {

    const {acctid} = req.params;
    var chaincode_id = 'assetboxcc';
    var function_name = 'getCoin';
    var arg_list = [acctid];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        if (result != null) {
            res.json({
                code: 200,
                coin: result
            });
        } else {
            res.json({
                code: 400,
                coin: null
            });
        }

    }, (err) => {
        console.error(err);
        res.status(500).json({
            code   : 500,
            message: 'server error',
        });
    });

});

router.get('/', function (req, res, next) {

    var chaincode_id = 'assetboxcc';
    var function_name = 'getAllAccount';
    var arg_list = [];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        if (result != null) {
            res.json({
                code: 200,
                coin: result
            });
        } else {
            res.json({
                code: 400,
                coin: null
            });
        }

    }, (err) => {
        console.error(err);
        res.status(500).json({
            code   : 500,
            message: 'server error',
        });
    });

});

module.exports = router;
