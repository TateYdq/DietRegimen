
var app = getApp();
const apiRequest = require("../../../utils/api_request.js")
const cache = require("../../../utils/cache.js")
Page({
  data: {
    foodArray: [],
    collectArray: [],
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.initData()
  },
  onShow: function () {
  },
  /**
 * 跳转到推荐文章
 */
  seeDetails: function (e) {
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
    apiRequest.getCollectFood(this.callbackGetFoodList)
  },
  callbackGetFoodList: function (res) {
    if (res.code == 2000) {
      this.setData({
        collectArray: res["collect_info"]
      });
      for (var i = 0; i < this.data.collectArray.length; i++) {
        var id = this.data.collectArray[i]['food_id']
        //提取文字
        var info = cache.getFoodInfo(id)
        if (info) {
          var param = {};
          var string = "foodArray[" + i + "]";
          param[string] = info;
          this.setData(param);
        }else{
          apiRequest.getFoodDetails(id,this.getDetailsCallback)
        }
        //提取图片
        var value = cache.getFoodImageValue(id)
        if (value) {
          var param = {};
          var string = "foodArray[" + i + "].localImagePath";
          param[string] = value;
          this.setData(param);
        } else {
          var path = this.data.foodArray[i]["photo_path"]
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
  getDetailsCallback: function(i,id,res){
    var param = {};
    var string = "foodArray[" + i + "]";
    param[string] = res["food_detail"];
    this.setData(param);
  },
  getImageCallback: function (i, id, res) {
    if (res.path) {
      var param = {};
      var string = "foodArray[" + i + "].localImagePath";
      param[string] = res.path;
      this.setData(param);
      cache.setFoodImage(id, res.path)
      console.log(this.data.foodArray[i].localImagePath)
    }
  }

});

