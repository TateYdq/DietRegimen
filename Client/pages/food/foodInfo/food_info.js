// pages/food/foodInfo/food_info.js
const static_data = require("../../../utils/static_data.js")
Page({

  /**
   * 页面的初始数据
   */
  data: {
      foodArray: [],
      foodInfo:{}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var foodID = options.id ;
    this.initData();
    var foodDetail = this.findFoodDetail(foodID);
    this.setData({
      "foodInfo.foodID":foodID,
      "foodInfo.name":foodDetail.name,
      "foodInfo.viewCount":foodDetail.viewCount,
      "foodInfo.foodKind":foodDetail.foodKind,
      "foodInfo.collectCount":foodDetail.collectCount,
      "foodInfo.info":foodDetail.info,
      "foodInfo.photoPath":foodDetail.photoPath,
    })
    console.log("name:%s", this.data.foodInfo['name']);
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
  initData: function(){
    this.setData({ foodArray: static_data.foodInfo });
  }, 
  findFoodDetail: function (foodID) {
    var i = 0;
    for (i = 0; i < this.data.foodArray.length; i++) {
      if (this.data.foodArray[i].foodID == foodID) {
        return this.data.foodArray[i];
      }
    }
    return null;
  }
})