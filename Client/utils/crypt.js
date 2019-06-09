
var fun_aes = require('../utils/aes.js')
var key = fun_aes.CryptoJS.enc.Utf8.parse("ydqydqydq");
//十六位十六进制数作为秘钥偏移量
var iv = fun_aes.CryptoJS.enc.Utf8.parse('1234567890654321');
function Encrypt(word){
  var srcs = fun_aes.CryptoJS.enc.Utf8.parse(word);
  var encrypted = fun_aes.CryptoJS.AES.encrypt(srcs, key, { iv: iv, mode: fun_aes.CryptoJS.mode.CBC, padding: fun_aes.CryptoJS.pad.Pkcs7 });
  return encrypted.ciphertext.toString().toUpperCase();
}
function Decrypt(word){
  var encryptedHexStr = fun_aes.CryptoJS.enc.Hex.parse(word);
  var srcs = fun_aes.CryptoJS.enc.Base64.stringify(encryptedHexStr);
  var decrypt = fun_aes.CryptoJS.AES.decrypt(srcs, key, { iv: iv, mode: fun_aes.CryptoJS.mode.CBC, padding: fun_aes.CryptoJS.pad.Pkcs7 });
  var decryptedStr = decrypt.toString(fun_aes.CryptoJS.enc.Utf8);
  return decryptedStr.toString();
}

module.exports = {
  Encrypt: Encrypt,
  Decrypt: Decrypt,
}