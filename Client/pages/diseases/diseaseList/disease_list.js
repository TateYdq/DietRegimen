// pages/diseases/diseases_list.js

const static_data = require("../../../utils/static_data.js")

Page({

  /**
   * 页面的初始数据
   */
  data: {
    diseasesArray:[],
    diseaArray:[
      {
        id: 1,
        photoPath: '../../../imgs/images/diseaseList1.jpg',
        kindName: '糖尿病',
        viewCount: 200,
        collectCount: 65,
        kindInfo: '重大突破！2016年糖尿病患者有望实现一年给药一次！'
      },
      {
        id: 2,
        photoPath: '../../../imgs/images/diseaseList2.jpg',
        kindName: '癌症治疗',
        viewCount: 220,
        collectCount: 69,
        kindInfo: '癌症治疗转变思路休眠法或成主流'
      }, {
        id: 3,
        photoPath: '../../../imgs/images/diseaseList3.jpg',
        kindName: '乳腺疾病',
        viewCount: 150,
        collectCount: 15,
        kindInfo: '如何预防乳腺疾病？常见措施有哪些'
      }, {
        id: 4,
        photoPath: '../../../imgs/images/diseaseList4.jpg',
        kindName: '湿疹',
        viewCount: 200,
        collectCount: 65,
        kindInfo: '湿疹越来越多见？六招教你有效预防湿疹'
      },
      {
        id: 5,
        photoPath: '../../../imgs/images/diseaseList5.jpg',
        kindName: '精神心理疾病',
        viewCount: 220,
        collectCount: 69,
        kindInfo: '为什么精神心理疾病患者日益增加？'
      }, {
        id: 6,
        photoPath: '../../../imgs/images/diseaseList6.jpg',
        kindName: 'HIV',
        viewCount: 150,
        collectCount: 15,
        kindInfo: '《柳叶刀》重磅：同性性行为HIV感染率可控制为0%'
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
 * 跳转到疾病文章
 */
  seeDiseaDetail: function (e) {
    var id = e.currentTarget.id;
    wx.navigateTo({
      url: '../diseaseInfo/disease_info?id=' + id,
    })
  },
  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  },
  initData: function(){
    var diseases = static_data.diseaseInfo
    this.setData({
      diseasesArray: diseases
    })
  }, 
  seeDetail: function (e) {
    var id = e.currentTarget.id;
    wx.navigateTo({
      url: '../diseaseInfo/disease_info?id=' + id,
    })
  }
})