/*
此页面为发起request请求页面
*/
const shema = 'https://'
const domain='api.ydq6.com'
const urlPrefix= shema + domain+'/DietRegimen/client'

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

//疾病相关URL
const GetDiseasesListsUrl = urlPrefix + '/health/getDiseasesLists'
const SearchDiseaseUrl = urlPrefix + '/health/searchDisease'

const GetDiseaseDetailsUrl= urlPrefix + '/health/getDiseaseDetails'
const CommentDiseaseUrl = urlPrefix + '/health/commentDisease'
const GetCommentDiseaseUrl = urlPrefix + '/health/getComment'

//推荐相关
const GetQuestionnaire = urlPrefix + '/recommend/getQuestionnaire'
const SubmitQuestionnaire = urlPrefix + '/recommend/submitQuestionnaire'
const GetRecInfo = urlPrefix + '/recommend/getRecInfo'

//文件相关
const GetImage = urlPrefix + '/file/getImage'
const GetVoice = urlPrefix + '/file/getVoice'

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
function getRequest(url,data){
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
      return res
    },
    fail(){
      return null
      console.log("error")
    }
  })
}
/*
  核心函数，post请求api,data以json格式传入
用法:
*/
function postRequest(data){
  var token = getToken()
  var rData = JSON.stringify(data)
  wx.request({
    url: GetUserInfoUrl,
    method: 'POST',
    header: {
      "Content-Type": "application/x-www-form-urlencoded",
      "token":token
    },
    data: rData,
    dataType: "json",
    success: function (subRes) {
      console.log(subRes.data)
      return subRes.data
    },
    fail: function () {
      console.log("请求失败")
    },
    complete: function () {
      console.log("完成")
    }
  })
}

function getUserInfo(){
  return getRequest(GetUserInfoUrl, null)
}

function getFoodCategory(){
  return getRequest(GetFoodCategoryUrl,null)
}
function getFoodDetails(foodID){
  var array = {
    "food_id":foodID
  }
  return getRequest(GetFoodDetailsUrl, array)
}

function getFoodListByKind(kind){
  var array = {
    "keyword":kind
  }
  return getRequest(SearchFoodUrl,array)
}

function getFoodListByKind(kind) {
  var array = {
    "keyword": kind
  }
  return getRequest(SearchFoodUrl, array)
}

function login(callback){
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
            callback(true)
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
  callback(false)
}

module.exports = {
  domain:domain,
  urlPrefix:urlPrefix,
  postRequest: postRequest,
  login:login,
  getUserInfo: getUserInfo,
  getFoodCategory: getFoodCategory,
  getFoodDetails: getFoodDetails,
  getFoodListByKind: getFoodListByKind
}