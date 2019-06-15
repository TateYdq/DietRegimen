// pages/mine/collectDisease/collect_disease.js
// pages/diseases/diseases_list.js

// const static_data = require("../../../utils/static_data.js")
const apiRequest = require("../../../utils/api_request.js")
const cache = require("../../../utils/cache.js")
Page({

  /**
   * 页面的初始数据
   */
  data: {
    diseaseArray: [],
    collectArray: [],
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
  seeDiseaseDetail: function (e) {
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
  initData: function () {
    apiRequest.getCollectDisease(this.callbackGetDiseasesLists)
  },
  callbackGetDiseasesLists: function (res) {
    if (res.code == 2000) {
      this.setData({
        collectArray: res["collect_info"]
      });
      for (var i = 0; i < this.data.collectArray.length; i++) {
        var id = this.data.collectArray[i]['disease_id']
        //提取文字
        var info = cache.getDiseaseInfo(id)
        if (info) {
          var param = {};
          var string = "diseaseArray[" + i + "]";
          param[string] = info;
          this.setData(param);
          //提取图片
          var value = cache.getDiseaseImageValue(id)
          if (value) {
            var params = {};
            var string2 = "diseaseArray[" + i + "].localImagePath";
            params[string2] = value;
            this.setData(params);
          } else {
            var path = this.data.diseaseArray[i]["photo_path"]
            apiRequest.getImage(i, id, path, this.getImageCallback)
          }
        } else {
          this.index_i = i
          this.index_id = id
          apiRequest.getDiseaseDetails(id, this.getDetailsCallback)
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
  getDetailsCallback: function (res) {
    var i = this.index_i
    var id = this.index_id
    var param = {};
    var string = "diseaseArray[" + i + "]";
    console.log("new:")
    console.log(res)
    param[string] = res["disease_detail"];
    this.setData(param);

    //提取图片
    var value = cache.getDiseaseImageValue(id)
    if (value) {
      var params = {};
      var string2 = "diseaseArray[" + i + "].localImagePath";
      params[string2] = value;
      this.setData(params);
    } else {
      var path = this.data.diseaseArray[i]["photo_path"]
      apiRequest.getImage(i, id, path, this.getImageCallback)
    }
  },
  getImageCallback: function (i, id, res) {
    if (res.path) {
      var param = {};
      var string = "diseaseArray[" + i + "].localImagePath";
      param[string] = res.path;
      this.setData(param);
      cache.setDiseaseImage(id, res.path)
      console.log(this.data.diseaseArray[i].localImagePath)
    }
  }
})