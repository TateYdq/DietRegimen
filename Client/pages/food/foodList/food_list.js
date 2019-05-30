// pages/food/foodList/food_list.js

const static_data = require("../../../utils/static_data.js")
const constant = require("../../../utils/constant.js")
Page({

  /**
   * 页面的初始数据
   */
  data: {
    foodArray:[],
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
      this.initData();
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
  /*
    初始化数据函数
  */
  initData: function(){
    this.setData({ foodArray: static_data.foodInfo });
  },
  seeDetail: function(){
    
  },
  switchNav: function(){
    wx.redirectTo({
      url: '../foodKind/food_kind',
    })
  }
})