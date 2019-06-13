var diseaseImageKeyStr = "disease_{id}_image" 
var diseaseInfoKeyStr = "disease_{id}_info"

var foodImageKeyStr = "food_{id}_image"
var foodInfoKeyStr = "food_{id}_info"
String.prototype.format = function (args) {
  var result = this;
  if (arguments.length < 1) {
    return result;
  }

  var data = arguments;       //如果模板参数是数组
  if (arguments.length == 1 && typeof (args) == "object") {
    //如果模板参数是对象
    data = args;
  }
  for (var key in data) {
    var value = data[key];
    if (undefined != value) {
      result = result.replace("{" + key + "}", value);
    }
  }
  return result;
}

function setFoodImage(id, value) {
  var foodImageKey = foodImageKeyStr.format({ 'id': id })
  try {
    wx.setStorageSync(foodImageKey, value)
  } catch (e) {
    console.log(e)
  }
}

function getFoodImageValue(id) {
  var foodImageKey = foodImageKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(foodImageKey)
    return value
  } catch (e) {
    console.log(e)
  }
}


function setDiseaseImage(id,value){
  var diseaseImageKey = diseaseImageKeyStr.format({'id':id})
  try {
    wx.setStorageSync(diseaseImageKey, value)
  } catch (e) { 
    console.log(e)
  }
}

function getDiseaseImageValue(id){
  var diseaseImageKey = diseaseImageKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(diseaseImageKey)
    return value
  } catch (e) {
    console.log(e)
    // Do something when catch error
  }
}

function setDiseaseInfo(id, value) {
  var diseaseInfoKey = diseaseInfoKeyStr.format({ 'id': id })
  try {
    var newvalue = JSON.stringify(value)
    wx.setStorageSync(diseaseInfoKey, newvalue)
  } catch (e) {
    console.log(e)
  }
}

function getDiseaseInfo(){
  var diseaseInfoKey = diseaseInfoKeyStr.format({ 'id': id })
  try {
    var newvalue = wx.getStorageSync(diseaseInfoKey)
    var value = JSON.parse(newvalue)
    return value
  } catch (e) {
    console.log(e)
    // Do something when catch error
  }
}

module.exports = {
  setFoodImage: setFoodImage,
  getFoodImageValue: getFoodImageValue,

  setDiseaseImage: setDiseaseImage,
  getDiseaseImageValue: getDiseaseImageValue,
  setDiseaseInfo: setDiseaseInfo,
  getDiseaseInfo: getDiseaseInfo,
}