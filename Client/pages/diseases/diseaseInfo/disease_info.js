// pages/diseases/diseasesInfo/diseases_info.js
const static_data = require("../../../utils/static_data.js")

Page({

  /**
   * 页面的初始数据
   */
  data: {
    diseasesArray: [],
    diseaseInfo: {}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var dID = options.id;
    this.initData();
    var dDetail = this.findDiseaseDetail(dID);
    this.setData({
      "diseaseInfo.diseaseID": dID,
      "diseaseInfo.name": dDetail.name,
      "diseaseInfo.viewCount": dDetail.viewCount,
      "diseaseInfo.diseaseKind": dDetail.diseaseKind,
      "diseaseInfo.collectCount": dDetail.collectCount,
      "diseaseInfo.info": dDetail.info,
      "diseaseInfo.photoPath": dDetail.photoPath,
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
  initData: function () {
    this.setData({ diseasesArray: static_data.diseaseInfo });
  },
  findDiseaseDetail: function (dID) {
    var i = 0;
    for (i = 0; i < this.data.diseasesArray.length; i++) {
      if (this.data.diseasesArray[i].diseaseID == dID) {
        return this.data.diseasesArray[i];
      }
    }
    return null;
  },
})