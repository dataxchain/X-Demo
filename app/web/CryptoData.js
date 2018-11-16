const crypto = require('crypto');


class CryptoData {

    static encrypt(readStream, writeStream, key) {

        let cipher = crypto.createCipher('aes-256-cbc', key);
        readStream.pipe(cipher).pipe(writeStream);
        writeStream.on('finish', function() {
            console.log('Encrypted file written!');
        });
    }

    static decrypt(readStream, writeStream, key) {

        let decipher = crypto.createDecipher('aes-256-cbc', key);
        readStream.pipe(decipher).pipe(writeStream);
        writeStream.on('finish', function() {
            console.log('Decrypted file written!');
        });
    }

    static generateRandomStringKey(number) {

        var ret = ""
        for (var i = 0; i < number; i++){
            ret += Math.random().toString(36).substr(2, 1)
        }
        return ret
    }
}

module.exports = CryptoData;

