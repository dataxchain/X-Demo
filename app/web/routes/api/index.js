const router = require('express').Router()
const auth = require('./v1')

router.use('/v1', auth)

module.exports = router