var diseaseImageKeyStr = "disease_{id}_image" 
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

function setDiseaseImage(id,value){
  var diseaseImageKey = diseaseImageKeyStr.format({'id':id})
  try {
    wx.setStorageSync(diseaseImageKey, value)
  } catch (e) { 
    console.log(e)
  }
}

function getDiseaseImageVlue(id){
  var diseaseImageKey = diseaseImageKeyStr.format({ 'id': id })
  try {
    var value = wx.getStorageSync(diseaseImageKey)
    return value
  } catch (e) {
    console.log(e)
    // Do something when catch error
  }
}


module.exports = {
  setDiseaseImage: setDiseaseImage,
  getDiseaseImageVlue: getDiseaseImageVlue,
}