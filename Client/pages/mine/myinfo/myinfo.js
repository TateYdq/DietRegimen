// pages/mine/myinfo/myinfo.js
const apiRequest = require("../../../utils/api_request.js")
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    myUserInfo: {},
    userNameRules: {
    maxLength: {
      value: 10,
      message: '姓名最多10个字',
    },
    minLength: {
      value: 1,
      message: '姓名最少1个字',
    },
    },
    isRequired: {
      required: {
        value: true,
        message: '必填',
      },
    },
    diseases_items: [
      {
        text: '高血压',
      },
      {
        text: '高血脂',
      },
      {
        text: '心脏病',
      },
      {
        text: '糖尿病',
      },
      {
        text: '老年痴呆症',
      },
      {
        text: '骨质疏松',
      },
      {
        text: '冠心病',
      },
      {
        text: '癌症',
      },
      {
        text: '风湿病',
      }
    ],
    gend_items:[
      {
        key: '男',
        value: 'male',
      },
      {
        key: '女',
        value: 'female'
      }
    ]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.initData()
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
  initData: function () {
    this.setData({
      myUserInfo: app.globalData.myUserInfo
    });
    // var diseasesFocus = '心脏病;高血压'
    this.diseaseFocus(this.data.myUserInfo['diseases_focus'])
  }, 
  diseaseFocus: function(diseasesFocus){
    var diseases = diseasesFocus.split(';')
    for(var i=0; i < diseases.length;i++){
      for (var j = 0; j < this.data.diseases_items.length;j++){
        if (this.data.diseases_items[j].text.indexOf(diseases[i]) != -1 || diseases[i].indexOf(this.data.diseases_items[j].text) != -1 ) {
          var params = {}
          var strings = "diseases_items["+j+"].checked"
          params[strings]=true
          this.setData(params)
          break
        }
      }
    }
  },
  handlePickerOpen(e) {
  },
  handlePickerCancel(e) {
    console.log(e)
  },
  wussFormSubmit(e) {
    console.log('提交了:', e.detail);
    var no_attenion = 0
    if(e.detail.isVip == false){
      no_attenion = 1
    }
    var diseasesFocusStr = ""
    for (var i = 0; i < e.detail.diseasesFocus.length;i++){
      diseasesFocusStr += e.detail.diseasesFocus[i].text+";"
    }
    apiRequest.updateUserInfo(e.detail.userName, e.detail.age, e.detail.gender, diseasesFocusStr, no_attenion,this.callbackSubmit)
  },
  wussFormReset(e) {
    console.log('重置了:', e.detail);
  },
  callbackSubmit: function(res){
      console.log(res.code)
    apiRequest.getUserInfo(this.callbackGetUserInfo)
  },
  callbackGetUserInfo: function (res) {
    console.log(res)
    if (res.code == 2000) {
      app.globalData.myUserInfo = res["user_info"]
      wx.navigateBack({
        delta: 1
      })
    }
  }
})