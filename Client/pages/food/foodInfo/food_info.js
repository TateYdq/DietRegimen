// pages/food/foodInfo/food_info.js
// const static_data = require("../../../utils/static_data.js")
const apiRequest = require("../../../utils/api_request.js")
const cache = require("../../../utils/cache.js")
Page({

  /**
   * 页面的初始数据
   */
  data: {
      foodArray: [],
      foodInfo:{},
      id: null,
      collected: false,
      localImagePath: null,
      localVoicePath: null,
      isPlay:false,
      currentPostId:null,
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var fID = Number(options.id)
    this.setData({
      id:fID
    })

    //收藏按钮、分享按钮、语言按钮---------------------------
    var postId = fID;
    this.setData({
      currentPostId: postId,
      isPlay:false
    })
    apiRequest.getFoodDetails(fID, this.callbackGetDetails)
    apiRequest.isCollectedFood(fID, this.isCollectCallback)
    this.registerAudioContext();
    // var postCollected = wx.getStorageSync('postCollected');
    // if (postCollected) { //取到缓存,设置相应的状态
    //   var collection = postCollected[postId];
    //   this.setData({ collected: collection });
    // } else {//没有取到缓存
    //   postCollected = {};
    //   postCollected[postId] = false;
    //   wx.setStorageSync('postCollected', postCollected);
    // }
  },
  registerAudioContext: function () {
    this.innerAudioContext = wx.createInnerAudioContext()
    this.innerAudioContext.src = this.data.localVoicePath
    this.innerAudioContext.onPlay(() => {
      console.log('开始播放')
    })
    this.innerAudioContext.onError((res) => {
      console.log(res.errMsg)
      console.log(res.errCode)
    })
    this.innerAudioContext.onEnded((res)=>{
      this.setData({ isPlay: false });
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
  // findFoodDetail: function (foodID) {
  //   var i = 0;
  //   for (i = 0; i < this.data.foodArray.length; i++) {
  //     if (this.data.foodArray[i].foodID == foodID) {
  //       return this.data.foodArray[i];
  //     }
  //   }
  //   return null;
  // },

  //---------------------------------------
  // 收藏btn
  onCollectedTap: function (event) {
    var collected = this.data.collected
    if (collected) {
      apiRequest.cancelCollectedFood(this.data.id, this.callbackCollect)
    } else {
      apiRequest.collectFood(this.data.id, this.callbackCollect)
    }

    // var that = this;

    // //异步获取缓存数据
    // wx.getStorage({
    //   key: 'postCollected',
    //   success: function (res) {
    //     var postCollected = res.data;
    //     var collection = postCollected[that.data.currentPostId];
    //     that.showToast(collection, postCollected);
    //   }
    // })
    // var collection = postCollected[this.data.currentPostId];
  },

  // showToast: function (collection, postCollected) {
  //   wx.showToast({
  //     title: collection ? "取消成功" : "收藏成功",
  //     icon: 'success',
  //     duration: 1000
  //   });
  //   collection = !collection;
  //   postCollected[this.data.currentPostId] = collection;
  //   //更新缓存
  //   wx.setStorageSync('postCollected', postCollected);
  //   //同步数据
  //   this.setData({ collected: collection });
  // },
  callbackGetDetails: function (res) {
    if (res.code == 2000) {
      this.setData({
        foodInfo: res["food_detail"]
      });
      cache.setFoodInfo(this.data.id,res["food_detail"])
      var value = cache.getFoodImageValue(this.data.id)
      if (value) {
        this.setData({
          localImagePath: value
        });
      } else {
        apiRequest.getImage(0, this.data.id, this.data.foodInfo["photo_path"], this.getImageCallback)
      }
      var file = cache.getFoodVoiceValue(this.data.id)
      if (file) {
        this.setData({
          localVoicePath: file
        });
      } else {
        apiRequest.downloadVoice(this.data.id, this.data.foodInfo["voice_path"], this.getVoiceCallback)
      }

    } else if (res.code == 4003) {

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
      this.setData({
        localImagePath: res.path
      });
      cache.setFoodImage(id, res.path)
    }
  },
  getVoiceCallback: function (id, res) {
    if (res.tempFilePath) {
      this.setData({
        localVoicePath: res.tempFilePath
      });
      cache.setFoodVoice(id, res.tempFilePath)
    }
  },
  // 分享btn
  onShareTap: function (event) {
    var itemList = ["分享到微信好友", "分享到朋友圈", "分享到QQ", "分享到微博"];

    wx.showActionSheet({
      itemList: itemList,
      itemColor: "#405f80",
      success: function (res) {

      }
    });
  },
  isCollectCallback: function (res) {
    if (res.code == 2000) {
      if (res.result == true) {
        this.setData({
          collected: true
        })
      }
    }
  },
  callbackCollect: function (res) {
    var collected = this.data.collected
    if (res.code == 2000) {
      wx.showToast({
        title: collected ? "取消成功" : "收藏成功",
        icon: 'success',
        duration: 1000
      })
      this.setData({
        collected: !collected,
      })
    } else if (res.code == 4003) {
      wx.showToast({
        title: '没有登录',
        icon: 'fail',
        duration: 1000
      })
    } else {
      wx.showToast({
        title: '失败',
        icon: 'fail',
        duration: 1000
      })
    }
  },
  // 播放语音btn
  onMusicTap: function (event) {
    this.innerAudioContext.src = this.data.localVoicePath
    var isPlayed = this.data.isPlay;//正在播放true
    if (isPlayed) {
      //暂停播放（或者关闭音乐：下次从头播放）
      //……
      this.innerAudioContext.pause()
      this.setData({ isPlay: false });
    } else {
      this.innerAudioContext.play()
      //开始播放
      //……
      this.setData({ isPlay: true });
    }
  }
})
