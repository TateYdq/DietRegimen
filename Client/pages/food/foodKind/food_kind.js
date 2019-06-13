// pages/food/foodKind/food_kind.js
var app = getApp();

const static_data = require("../../../utils/static_data.js")
const apiRequest = require("../../../utils/api_request.js")

Page({
  data: {
    foodKindArray: [],
    foodArray:[
      {
        id:1,
        photoPath: '../../../imgs/images/tuiList1.jpg',
        kindName: '养生菇粥',
        viewCount: 200,
        collectCount:65,
        kindInfo: '功效：安定凝神、利尿通便、健胃消食。'
      },
      {
        id: 2,
        photoPath: '../../../imgs/images/tuiList2.jpg',
        kindName: '养生汤品',
        viewCount: 220,
        collectCount: 69,
        kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      }, {
        id: 3,
        photoPath: '../../../imgs/images/tuiList3.jpg',
        kindName: '八宝养生饭',
        viewCount: 150,
        collectCount: 15,
        kindInfo: '功效：减肥排毒、润肺止咳、安定凝神。'
      }
    ], 
    newsArray: [
      {
        id: 1,
        photoPath: '../../../imgs/images/newsList1.jpg',
        kindName: '养生锅做法',
        viewCount: 200,
        collectCount: 65,
        kindInfo: '做法：1.羊肉洗净，切块；虾子，洗净；山药洗净，切片备用。'
      },
      {
        id: 2,
        photoPath: '../../../imgs/images/newsList2.jpg',
        kindName: '食谱名2',
        viewCount: 220,
        collectCount: 69,
        kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      }, {
        id: 3,
        photoPath: '../../../imgs/images/newsList3.jpg',
        kindName: '红芹菜的功效',
        viewCount: 150,
        collectCount: 15,
        kindInfo: '平肝降压、镇静安神、多食芹菜有利于安定情绪，消除烦躁。防癌抗癌。'
      }
    ],
    banners: [
      {
        id: 3,
        img: '../../../imgs/images/homeBig1.jpg',
        url: '',
        name: '轮播图1'
      },
      {
        id: 1,
        img: '../../../imgs/images/homeBig2.jpg',
        url: '',
        name: '轮播图2'
      },
      {
        id: 2,
        img: '../../../imgs/images/homeBig3.jpg',
        url: '',
        name: '轮播图3'
      }
    ]
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
      var array = this.initData();
      this.setData({foodKindArray:array});
  },
  onShow: function () {
  }, 
    /**
   * 跳转到推荐文章
   */
  seeRecDetail: function (e) {
    var id = e.currentTarget.id;
    wx.navigateTo({
      url: '../foodInfo/food_info?id=' + id,
    })
  },
  seeNewsDetail: function(e) {
    var id = e.currentTarget.id;
    wx.navigateTo({
      url: '../foodInfo/food_info?id=' + id,
    })
  },
    /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  },
  initData: function () {
    return static_data.foodKindInfo;
  },
  seeFoodLists: function (e) {
    var kindName = e.currentTarget.dataset.name;
    console.log("kindName:%s",kindName);
    wx.navigateTo({
      url: '../foodList/food_list?kindName=' + kindName+'&fromKind=true',
    })
  }
});
//   /**
//    * 生命周期函数--监听页面初次渲染完成
//    */
//   onReady: function () {

//   },

//   /**
//    * 生命周期函数--监听页面显示
//    */
//   onShow: function () {

//   },

//   /**
//    * 生命周期函数--监听页面隐藏
//    */
//   onHide: function () {

//   },

//   /**
//    * 生命周期函数--监听页面卸载
//    */
//   onUnload: function () {

//   },

//   /**
//    * 页面相关事件处理函数--监听用户下拉动作
//    */
//   onPullDownRefresh: function () {

//   },

//   /**
//    * 页面上拉触底事件的处理函数
//    */
//   onReachBottom: function () {

//   },
