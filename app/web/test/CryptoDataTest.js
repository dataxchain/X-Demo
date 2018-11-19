var CryptoData = require('../CryptoData');
var fs = require('fs');
var assert = require("assert"); //nodejs에서 제공하는 aseert 모듈

describe('CryptoFile test!', function () {
    it('generate random string test', function () {
        for (var i = 0 ; i < 100000 ; i ++){
            assert.equal(32, CryptoData.generateRandomStringKey(32).length);
        }
    });

    it('text encrypt and decrypt test', function () {
        let key = 'a1b2c3d4';

        let readStream = fs.createReadStream('text.txt');
        let writeStream = fs.createWriteStream('text.txt.enc');
        CryptoData.encrypt(readStream, writeStream, key);
    });

    it('text encrypt and decrypt test', function () {
        let key = 'a1b2c3d4';
        let encryptedReadStream = fs.createReadStream('text.txt.enc');
        let encryptedwriteStream = fs.createWriteStream('result.txt');
        CryptoData.decrypt(encryptedReadStream, encryptedwriteStream, key);

    });

    describe('file encrypt and decrypt test!', function () {
        const key = '9wb16j4zcvncytz76hq826q96s9733gs'
        const fileName = 'KindleForMac-50131.dmg';

        it('file encrypt test', function () {

            let readStream = fs.createReadStream(fileName);
            let writeStream = fs.createWriteStream(fileName + '.enc');
            CryptoData.encrypt(readStream, writeStream, key);

        });

        it('file decrypt test', function () {

            var encryptedFileStream = fs.createReadStream(fileName + '.enc');
            var decryptedFileStream = fs.createWriteStream('result.dmg');
            CryptoData.decrypt(encryptedFileStream, decryptedFileStream, key);

        });
    });
});