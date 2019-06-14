/*
此页面为发起request请求页面
*/
const shema = 'https://'
const domain='api.ydq6.com'
const urlPrefix= shema + domain+'/DietRegimen/client'

const filePrefix = shema + domain +'/DietRegimen/file'
const SUCCESS = 2000
const Forbidden = 4003
const Failed = 5000
const ServerError = 5003
//用户相关url
const UserLoginUrl = urlPrefix + '/user/userLogin'

const GetUserInfoUrl = urlPrefix+'/user/getUserInfo'
const UpdateUserInfoUrl = urlPrefix + '/user/updateUserInfo'

const CollectFoodUrl = urlPrefix +'/user/collectFood'
const GetCollectFoodUrl = urlPrefix + '/user/getCollectFood'

const CollectDiseaseUrl = urlPrefix + '/user/collectDisease'
const GetCollectDiseaseUrl = urlPrefix + '/user/getCollectDisease'

//食物相关URL
const GetFoodCategoryUrl = urlPrefix + '/food/getFoodCategory'
const SearchFoodUrl = urlPrefix + '/food/searchFood'

const GetFoodDetailsUrl = urlPrefix + '/food/getFoodDetails'
const CommentFoodUrl = urlPrefix + '/food/commentFood'
const GetCommentFoodUrl = urlPrefix + '/food/getComment'

const IsCollectedFoodUrl = urlPrefix + '/food/isCollected'
const CancelCollectedFoodUrl = urlPrefix + '/food/cancelCollected'

//疾病相关URL
const GetDiseasesListsUrl = urlPrefix + '/health/getDiseasesLists'
const SearchDiseaseUrl = urlPrefix + '/health/searchDisease'

const GetDiseaseDetailsUrl= urlPrefix + '/health/getDiseaseDetails'
const CommentDiseaseUrl = urlPrefix + '/health/commentDisease'
const GetCommentDiseaseUrl = urlPrefix + '/health/getComment'

const IsCollectedDiseaseUrl = urlPrefix + '/health/isCollected'
const CancelCollectedDiseaseUrl = urlPrefix + '/health/cancelCollected'

//推荐相关
const GetQuestionnaire = urlPrefix + '/recommend/getQuestionnaire'
const SubmitQuestionnaire = urlPrefix + '/recommend/submitQuestionnaire'
const GetRecInfoUrl = urlPrefix + '/recommend/getRecInfo'

//文件相关
const GetImageUrl = filePrefix + '/getImage'
const GetVoiceUrl = filePrefix + '/getVoice'

const crypt = require("../utils/crypt.js")

const app = getApp()

const promisify = require('./promisify.js')
const myRequest = promisify(wx.request)
const myLogin = promisify(wx.login)

function getToken(){
  try{
    var token = wx.getStorageSync("token")
    // console.log("解密前:", token)
    // token = crypt.Decrypt(token)
    // console.log("解密后:",token)
    return token
  }catch(e){
    console.log("未登录")
    return ""
  }
}
/*
核心函数，get请求api,data以json格式传入
用法:
*/
function getRequest(url, data, callback){
  var token = getToken()
  wx.request({
    url:url,
    method:"GET",
    header: {
      "token": token
    },
    data:data,
    success(res){
      console.log("res:\n")
      console.log(res.data)
      callback(res.data)
    },
    fail(){
      console.log("error")
    }
  })
}
/*
  核心函数，post请求api,data以json格式传入
用法:
*/
function postRequest(url,data,callback){
  var token = getToken()
  var rData = JSON.stringify(data)
  wx.request({
    url: url,
    method: 'POST',
    header: {
      "Content-Type": "application/x-www-form-urlencoded",
      "token":token
    },
    data: rData,
    dataType: "json",
    success: function (res) {
      console.log(res.data)
      callback(res.data)
    },
    fail: function () {
      console.log("请求失败")
    },
    complete: function () {
      console.log("完成")
    }
  })
}

function getUserInfo(callback){
  getRequest(GetUserInfoUrl, null, callback)
}
function isCollectedFood(foodID, callback) {
  var array = {
    "food_id": foodID
  }
  getRequest(IsCollectedFoodUrl, array, callback)
}

function isCollectedDisease(diseaseID,callback){
  var array = {
    "disease_id": diseaseID
  }
  getRequest(IsCollectedDiseaseUrl, array,callback)
}

function getFoodCategory(callback){
  getRequest(GetFoodCategoryUrl, null, callback)
}
function getFoodDetails(foodID, callback){
  var array = {
    "food_id":foodID
  }
  getRequest(GetFoodDetailsUrl, array, callback)
}

function getFoodListByKind(kind, callback){
  var array = {
    "keyword":kind
  }
  getRequest(SearchFoodUrl, array, callback)
}

function getFoodList(callback) {
  getRequest(SearchFoodUrl, null,callback)
}
function getDiseasesLists(callback){
  getRequest(GetDiseasesListsUrl, null, callback)
}

function getDiseaseDetails(diseaseID,callback){
  var data = {
    "disease_id": diseaseID
  }
  getRequest(GetDiseaseDetailsUrl, data, callback)
}

function getRecInfo(callback){
  getRequest(GetRecInfoUrl, null, callback)
}


function updateUserInfo(name, age, gender, diseasesFocus, callback){
  var userInfo = {
    "user_info":{
      "name":name,
      "age":age,
      "gender":gender,
      "diseases_focus": diseasesFocus
    }
  }
  postRequest(UpdateUserInfoUrl,userInfo,callback)

}

function collectDisease(diseaseID, callback){
  var data ={
    "disease_id":diseaseID
  }
  postRequest(CollectDiseaseUrl, data, callback)
}

function collectFood(foodID, callback) {
  var data = {
    "food_id": foodID
  }
  postRequest(CollectFoodUrl, data, callback)
}

function cancelCollectedFood(foodID, callback) {
  var data = {
    "food_id": foodID
  }
  getRequest(CancelCollectedFoodUrl, data, callback)
}

function cancelCollectedDisease(diseaseID, callback) {
  var data = {
    "disease_id": diseaseID
  }
  getRequest(CancelCollectedDiseaseUrl, data, callback)
}

function getImage(i,id,photoPath,callback){
  if(photoPath){
    console.log("url:"+GetImageUrl + "?path=" + photoPath)
    wx.getImageInfo({
      src: GetImageUrl+"?path="+photoPath,
      success(res) {
        callback(i,id,res)
      },
      fail(res) {
        callback(i,id,res)
      }
    })
  }
}

function login(callback){
  var result = new Array()
  result.code = 5003
  console.log("url:" + app.globalData.userInfo.avatarUrl)
  myLogin().then(res=>{
    console.log("返回码: " + res.code)
    if (res.code) {
      var rData = JSON.stringify({ 
        code: res.code,
        nick_name: app.globalData.userInfo.nickName,
        avatar_url: app.globalData.userInfo.avatarUrl,
        gender: app.globalData.userInfo.gender=2?"female":"male"
      })
      myRequest({
        url: UserLoginUrl,
        method: 'POST',
        header: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        data: rData,
        dataType: "json",
      }).then(subRes=>{
        console.log(subRes.data.token)
        if (subRes.data.code == SUCCESS) {
          try {
            var token = subRes.data.token
            // console.log("加密前:", subRes.data.token)
            // token = crypt.Encrypt(token)
            // console.log("加密后:",token)
            wx.setStorageSync('token', token)
            console.log("存储成功")
            app.globalData.isLogin = true
            wx.showToast({
              title: '登录成功',
              icon: 'success',
              image: '',
              duration: 2000,
              mask: false,
            })
            callback(subRes.data)
          } catch (e) {
            console.log("token存储失败")

          }
        }
      }).catch(subRes=>{
        console.log("请求失败")
      })
    }else{
      console.log("登录失败")
    }
  }).catch(res=>{
    console.log("连接失败")
  })
}

module.exports = {
  domain:domain,
  urlPrefix:urlPrefix,
  postRequest: postRequest,
  login:login,
  getUserInfo: getUserInfo,
  getFoodCategory: getFoodCategory,
  getFoodDetails: getFoodDetails,
  getFoodList: getFoodList,
  getDiseasesLists: getDiseasesLists,
  getDiseaseDetails: getDiseaseDetails,
  getImage: getImage,
  updateUserInfo: updateUserInfo,
  collectDisease: collectDisease,
  collectFood: collectFood,
  isCollectedDisease: isCollectedDisease,
  isCollectedFood: isCollectedFood,
  cancelCollectedFood: cancelCollectedFood,
  cancelCollectedDisease: cancelCollectedDisease,

  getRecInfo: getRecInfo,
}