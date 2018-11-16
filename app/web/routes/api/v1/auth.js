var express = require('express');
var router = express.Router();

var queryCC = require('../../../../middleware/query_chaincode.js');
//var invokeCC = require('../../../../middleware/invoke_chaincode.js');

router.post('/login', function(req, res, next) {

    const {account, password} = req.body;
    console.log(account, password);
    var chaincode_id = 'assetboxcc';
    var function_name = 'loginAccount';
    var arg_list = [account, password];

    queryCC.query_chaincode(chaincode_id, function_name, arg_list).then((result) => {
        if (result != null) {
            res.json({
                code: 200,
                account: result
            });
        }else {
            res.json({
                code: 401,
                message: "login failed"
            });
        }

    }, (err) => {
        console.error(error);
        res.status(500).json({
            code: 500,
            message: 'server error',
        });
    });

});

module.exports = router;
