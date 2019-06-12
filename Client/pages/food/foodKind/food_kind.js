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
        kindName: '食谱名1',
        viewCount: 200,
        collectCount:65,
        kindInfo:'功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      },
      {
        id: 2,
        photoPath: '../../../imgs/images/tuiList2.jpg',
        kindName: '食谱名2',
        viewCount: 220,
        collectCount: 69,
        kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      }, {
        id: 3,
        photoPath: '../../../imgs/images/tuiList3.jpg',
        kindName: '食谱名3',
        viewCount: 150,
        collectCount: 15,
        kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
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
    ],
    // shops: [
    //   {
    //     id: 1,
    //     img: '../../../imgs/images/tuiList1.jpg',
    //     distance: 1.8,
    //     sales: 1475,
    //     logo: 'http://wxapp.im20.com.cn/impublic/waimai/imgs/shops/logo_1.jpg',
    //     name: '杨国福麻辣烫(东四店)',
    //     desc: '满25减8；满35减10；满60减15（在线支付专享）'
    //   },
    //   {
    //     id: 2,
    //     img: '../../../imgs/images/tuiList2.jpg',
    //     distance: 2.4,
    //     sales: 1284,
    //     logo: '../../../imgs/images/tuiList3.jpg',
    //     name: '忠友麻辣烫(东四店)',
    //     desc: '满25减8；满35减10；满60减15（在线支付专享）'
    //   },
    //   {
    //     id: 3,
    //     img: '../../../imgs/images/tuiList4.jpg',
    //     distance: 2.3,
    //     sales: 2039,
    //     logo: 'http://wxapp.im20.com.cn/impublic/waimai/imgs/shops/logo_3.jpg',
    //     name: '粥面故事(东大桥店)',
    //     desc: '满25减8；满35减10；满60减15（在线支付专享）'
    //   }
    // ]
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
