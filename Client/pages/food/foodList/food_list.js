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
    console.log("fromkind:%s,kindName:%s", options.fromKind, options.kindName);
    var fromKind = options.fromKind
    if (fromKind && fromKind == "true"){
      var kindName = options.kindName;
      this.initData(true,kindName);
    }else{
      this.initData(false,"");
    }
    
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
  initData: function(fromKind,kindName){
    // console.log("fromkind:%s,kindName:%s", fromKind,kindName);
    if(fromKind){
      var i,j;
      j = 0;
      var arrays = new Array()
      for(i = 0;i < static_data.foodInfo.length;i++){
        // console.log("static_data.kindName:%s,kindName:%s", static_data.foodInfo[i].foodKind, kindName);
        if(static_data.foodInfo[i].foodKind == kindName){
          arrays[j] = static_data.foodInfo[i];
          j++;
        }
      }
      this.setData({
        "foodArray": arrays
      });
    }else{
      this.setData({ 
        foodArray: static_data.foodInfo 
      });
    }
  },
  seeDetail: function(e){
    var id = e.currentTarget.id;
    wx.navigateTo({
      url: '../foodInfo/food_info?id=' + id,
    })
  },
  switchNav: function(){
    wx.redirectTo({
      url: '../foodKind/food_kind',
    })
  },
  findFoodDetail: function(foodID){
      var i = 0;
      for(i = 0;i < foodArray.length;i++){
          if(foodArray[i].foodID == foodID){
            return foodArray[i];
          }
      }
      return null;
  }
})