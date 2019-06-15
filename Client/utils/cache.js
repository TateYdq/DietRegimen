var diseaseImageKeyStr = "disease_{id}_image" 
var diseaseInfoKeyStr = "disease_{id}_info"
var diseaseVoiceKeyStr = "disease_{id}_voice"

var foodImageKeyStr = "food_{id}_image"
var foodInfoKeyStr = "food_{id}_info"
var foodVoiceKeyStr = "food_{id}_voice"


var questionInfoKyeStr = "question_{id}_info"
var questionAnswerdKeyStr = "question_{id}_ans"
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

function isQuestionAnswered(id){
  console.log("clear")
  var key = questionAnswerdKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(key)
    if(value){
      return true
    }else{
      return false
    }
  } catch (e) {
    console.log(e)
    return false
  }
}

function clearAllQuestion(){
  for(var i = 1;i < 100;i++){
    var key = questionAnswerdKeyStr.format({ 'id': i })
    // console.log("key:"+key)
    wx.removeStorage({
      key: key,
      // success:function(res){
      //     console.log(res)
      // }
    })
  }
}

function removeAnswer(id){
  var key = questionAnswerdKeyStr.format({ 'id': id })
  try {
    wx.removeStorageSync(key)
  } catch (e) {
    console.log(e)
  }
}

function answerQuestion(id) {
  var key = questionAnswerdKeyStr.format({ 'id': id })
  try {
    wx.setStorageSync(key, "true")
  } catch (e) {
    console.log(e)
  }
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

function setDiseaseVoice(id, value) {
  var diseaseVoiceKey = diseaseVoiceKeyStr.format({ 'id': id })
  try {
    wx.setStorageSync(diseaseVoiceKey, value)
  } catch (e) {
    console.log(e)
  }
}

function getDiseaseVoiceValue(id) {
  var diseaseVoiceKey = diseaseVoiceKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(diseaseVoiceKey)
    return value
  } catch (e) {
    console.log(e)
    // Do something when catch error
  }
}

function setFoodVoice(id, value) {
  var foodVoiceKey = foodVoiceKeyStr.format({ 'id': id })
  try {
    wx.setStorageSync(foodVoiceKey, value)
  } catch (e) {
    console.log(e)
  }
}

function getFoodVoiceValue(id) {
  var foodVoiceKey = foodVoiceKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(foodVoiceKey)
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

function getDiseaseInfo(id){
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

function setFoodInfo(id, value) {
  var foodInfoKey = foodInfoKeyStr.format({ 'id': id })
  try {
    var newvalue = JSON.stringify(value)
    wx.setStorageSync(foodInfoKey, newvalue)
  } catch (e) {
    console.log(e)
  }
}

function getFoodInfo(id) {
  var foodInfoKey = foodInfoKeyStr.format({ 'id': id })
  try {
    var newvalue = wx.getStorageSync(foodInfoKey)
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
  setFoodInfo: setFoodInfo,
  getFoodInfo: getFoodInfo,

  setDiseaseVoice: setDiseaseVoice,
  getDiseaseVoiceValue: getDiseaseVoiceValue,
  setFoodVoice: setFoodVoice,
  getFoodVoiceValue: getFoodVoiceValue,

  setDiseaseImage: setDiseaseImage,
  getDiseaseImageValue: getDiseaseImageValue,
  setDiseaseInfo: setDiseaseInfo,
  getDiseaseInfo: getDiseaseInfo,

  isQuestionAnswered: isQuestionAnswered,
  answerQuestion: answerQuestion,
  removeAnswer: removeAnswer,
  clearAllQuestion: clearAllQuestion,
}