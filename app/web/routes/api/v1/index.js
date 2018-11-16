var express = require('express');
var router = express.Router();

var assetboxRouter = require('./assetbox');
var accountsRouter = require('./accounts');
var authRouter = require('./auth');


router.use('/assetbox', assetboxRouter);
router.use('/accounts', accountsRouter);
router.use('/auth', authRouter);

module.exports = router;
