// pages/common/common.js

const apiRequest = require("../../utils/api_request.js")
Page({

  /**
   * 页面的初始数据
   */
  data: {
    id:null,
    object:null,
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
      var id = Number(options.id)
      var object = options.object
      this.setData({
        id:id,
        object:object
      })
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
  formSubmit: function(e){
    var content = e.detail.value.content
    if(this.data.object == "food"){
      apiRequest.commentFood(this.data.id, content, this.submitCallback)
    } else if (this.data.object == "disease"){
      apiRequest.commentDisease(this.data.id, content, this.submitCallback)
    }
  },
  submitCallback: function(res){
    if(res.code==2000){
      wx.showToast({
        title: '评论成功',
        icon: 'success',
        duration: 2000,
      })
      wx.navigateBack({
        delta: 1
      })
    }else{
      wx.showToast({
        title: '评论失败',
        icon: 'success',
        duration: 1000,
      })
    }
  }
})