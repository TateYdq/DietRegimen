// pages/mine/mycenter/mycenter.js
const apiRequest = require("../../../utils/api_request.js")
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
      isLogin:false,
      userAvatarUrl: null,
      myUserInfo:{}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.setData({
      isLogin: app.globalData.isLogin
    });
    console.log(app.globalData.isLogin)
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
    this.setData({
      myUserInfo: app.globalData.myUserInfo
    });
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  },
  getWechatUserInfo: function (e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      myUserInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },
  setMyInfo: function(){
    wx.navigateTo({
      url: '../myinfo/myinfo',
    })
  },
  gotoCollectFood: function() {
    
    wx.navigateTo({
      /*
      url: '../collectFood/collect_food',
      暂时跳转到食物清单页
      */
      url: '../../food/foodList/food_list'
    })
  },
  
  gotoCollectDisease: function() {
    /*
    wx.navigateTo({
      url: '../collectDisease/collect_disease',
    })
     暂时跳转到疾病清单页
    */
    wx.switchTab({
      url: '../../diseases/diseaseList/disease_list',
    })
  },
  gotoFoodList: function(){
    wx.navigateTo({
      url: '../../food/foodList/food_list',
    })
  }, 
  login: function (e) {
    app.globalData.userInfo = e.detail.userInfo
    apiRequest.login(this.callbackLogin)
  },
  callbackLogin: function(res){
    if(res.code == 2000){
      this.setData({
        isLogin: true,
        userAvatarUrl: app.globalData.userInfo.avatarUrl
      });
      apiRequest.getUserInfo(this.callbackGetUserInfo)
    }
  },
  callbackGetUserInfo: function(res){
    console.log(res)
    if (res.code == 2000) {
      app.globalData.myUserInfo = res["user_info"]
      this.setData({
        myUserInfo: app.globalData.myUserInfo
      });
    } else if (res.code == 4003) {
      wx.showToast({
        title: '没有登录',
        icon: 'fail',
        image: '',
        duration: 2000,
        mask: false,
      })
    } else{
      wx.showToast({
        title: '失败',
        icon: 'fail',
        image: '',
        duration: 2000,
        mask: false,
      })
    }
}
})