// pages/food/foodKind/food_kind.js
var app = getApp();

// const static_data = require("../../../utils/static_data.js")
const apiRequest = require("../../../utils/api_request.js")
const cache = require("../../../utils/cache.js")
Page({
  data: {
    isLogin:false,
    isRec:false,
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
    newsArray:[],
    // newsArray: [
    //   {
    //     id: 1,
    //     photoPath: '../../../imgs/images/newsList1.jpg',
    //     kindName: '养生锅做法',
    //     viewCount: 200,
    //     collectCount: 65,
    //     kindInfo: '做法：1.羊肉洗净，切块；虾子，洗净；山药洗净，切片备用。'
    //   },
    //   {
    //     id: 2,
    //     photoPath: '../../../imgs/images/newsList2.jpg',
    //     kindName: '食谱名2',
    //     viewCount: 220,
    //     collectCount: 69,
    //     kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
    //   }, {
    //     id: 3,
    //     photoPath: '../../../imgs/images/newsList3.jpg',
    //     kindName: '红芹菜的功效',
    //     viewCount: 150,
    //     collectCount: 15,
    //     kindInfo: '平肝降压、镇静安神、多食芹菜有利于安定情绪，消除烦躁。防癌抗癌。'
    //   }
    // ],
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
      this.setData({
        isLogin: app.globalData.isLogin
      });
      this.initData()
  },
  onShow: function () {
    this.setData({
      isLogin: app.globalData.isLogin
    });
    //登录后没有推荐
    if(this.data.isLogin==true&&this.data.isRec==false){
      apiRequest.getRecInfo(this.callbackGetRec)
    }
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
    apiRequest.getFoodList(this.callbackGetFoodList)
  },
  callbackGetFoodList: function (res) {
    if (res.code == 2000) {
      this.setData({
        newsArray: res["food_list"]
      });
      var i;
      for (i = 0; i < this.data.newsArray.length; i++) {
        var id = this.data.newsArray[i]['food_id']
        // cache.setFoodInfo(id, this.data.newsArray[i])
        var value = cache.getFoodImageValue(id)
        if (value) {
          var param = {};
          var string = "newsArray[" + i + "].localImagePath";
          param[string] = value;
          this.setData(param);
        } else {
          var path = this.data.newsArray[i]["photo_path"]
          apiRequest.getImage(i, id, path, this.getImageCallback)
        }
      }

    } else if (res.code == 4003) {
      wx.showToast({
        title: '没有登录',
        icon: 'fail',
        image: '',
        duration: 2000,
        mask: false,
      })
    } else {
      wx.showToast({
        title: '失败',
        icon: 'fail',
        image: '',
        duration: 2000,
        mask: false,
      })
    }
  },
  getImageCallback: function (i, id, res) {
    if (res.path) {
      var param = {};
      var string = "newsArray[" + i + "].localImagePath";
      param[string] = res.path;
      this.setData(param);
      cache.setFoodImage(id, res.path)
      console.log(this.data.newsArray[i].localImagePath)
    }
  },
  callbackGetRec: function(res){
    if(res.code==2000){
      this.setData({
        foodArray: res["food_list"],
        isRec:true
      })
    }
  },
  seeFoodLists: function (e) {
    var kindName = e.currentTarget.dataset.name;
    console.log("kindName:%s",kindName);
    wx.navigateTo({
      url: '../foodList/food_list?kindName=' + kindName+'&fromKind=true',
    })
  }
});

