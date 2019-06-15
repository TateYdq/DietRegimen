// pages/diseases/diseasesInfo/diseases_info.js
// const static_data = require("../../../utils/static_data.js")
const apiRequest = require("../../../utils/api_request.js")
const cache = require("../../../utils/cache.js")

Page({

  /**
   * 页面的初始数据
   */
  data: {
    id: null,
    diseaseInfo: {},
    collected: false,
    localImagePath: null,
// <<<<<<< HEAD
    commonList: [
      {
        id: 1,
        userPhoto: '../../../imgs/images/author.jpg',
        userName: '曹总',
        contant: '这篇文章是我写的！',
        lytime: '2019-6-15 18:19:57',
        replyUserName: ""
        // collectCount: 65,
        // kindInfo: '功效：安定凝神、利尿通便、健胃消食。'
      },
      {
        id: 2,
        userPhoto: '../../../imgs/images/comment2.jpg',
        userName: 'CFO',
        contant: '曹总的文章写的真好，夸!',
        lytime: '2019-6-15 18:19:57',
        replyUserName: ''
        // collectCount: 69,
        // kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      }, {
        id: 3,
        userPhoto: '../../../imgs/images/comment3.jpg',
        userName: '王老板',
        contant: '曹总辛苦了!',
        lytime: '2019-6-15 18:19:57',
        replyUserName: ''
        // collectCount: 15,
        // kindInfo: '功效：减肥排毒、润肺止咳、安定凝神。'
      }, {
        id: 4,
        userPhoto: '../../../imgs/images/comment4.jpg',
        userName: '余博士',
        contant: '国际交互设计集团的大家都辛苦了！',
        lytime: '2019-6-15 18:19:57'
        // collectCount: 15,
        // kindInfo: '功效：减肥排毒、润肺止咳、安定凝神。'
      }
    ], 
// =======
    localVoicePath: null,
    isPlay: false,
    currentPostId: null,
  },
  // registerAudioContext: function () {
  //   this.innerAudioContext = wx.createInnerAudioContext()
  //   this.innerAudioContext.src = this.data.localVoicePath
  //   this.innerAudioContext.onPlay(() => {
  //     console.log('开始播放')
  //   })
  //   this.innerAudioContext.onError((res) => {
  //     console.log(res.errMsg)
  //     console.log(res.errCode)
  //   })
  //   this.innerAudioContext.onEnded((res)=>{
  //     this.setData({ isPlay: false });
  //   })
// >>>>>>> 32fe8efdf1dae0684edbcd504747d07c918e1712
  //},

  /**
  * 跳转到评论页
  */
  writeComment: function (e) {
    if (app.globalData.isLogin == false) {
      wx.showToast({
        title: '请先登录',
        icon: "none",
        duration: 1000
      })
      return
    }
    var id = this.data.id;
    wx.navigateTo({
      url: '../../common/common?id=' + id + '&object=disease',
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var dID = Number(options.id);
    this.setData({
      id: dID
    })

    
    apiRequest.getDiseaseDetails(dID, this.callbackGetDetails)
    //收藏按钮、分享按钮、语言按钮---------------------------
    var postId = dID;
    this.setData({
      currentPostId: postId,
      isPlay: false
    })
    apiRequest.isCollectedDisease(dID, this.isCollectCallback)
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
    this.innerAudioContext.onEnded((res) => {
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
    apiRequest.getFoodComment(this.data.id, this.getCommentCallback)
  },
  getCommentCallback: function (res) {
    if (res.code == 2000) {
      console.log("get comment res:")
      console.log(res)
      this.setData({
        commonList: res['comment_list']
      })
    }
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
  callbackGetDetails: function(res){
    if (res.code == 2000) {
      this.setData({
        diseaseInfo: res["disease_detail"]
      });
      cache.setDiseaseInfo(this.data.id, res["disease_detail"])
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
    var value = cache.getDiseaseImageValue(this.data.id)
    if (value) {
      this.setData({
        localImagePath: value
      });
    } else {
      apiRequest.getImage(0,this.data.id, this.data.diseaseInfo["photo_path"], this.getImageCallback)
    }
    var file = cache.getDiseaseVoiceValue(this.data.id)
    if (file) {
      this.setData({
        localVoicePath: file
      });
    } else {
      apiRequest.downloadVoice(this.data.id, this.data.diseaseInfo["voice_path"], this.getVoiceCallback)
    }
  },
  getImageCallback: function (id, res) {
    if (res.path) {
      this.setData({
        localImagePath: res.path
      });
      cache.setDiseaseImage(id,res.path)
    }
  },
  getVoiceCallback: function (id, res) {
    if (res.tempFilePath) {
      this.setData({
        localVoicePath: res.tempFilePath
      });
      cache.setDiseaseVoice(id, res.tempFilePath)
    }
  },

  //---------------------------------------
  // 收藏btn
  onCollectedTap: function (event){
    var collected = this.data.collected
    if (collected){
      apiRequest.cancelCollectedDisease(this.data.id, this.callbackCollect)
    }else{
      apiRequest.collectDisease(this.data.id, this.callbackCollect)
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
  },
  callbackCollect: function(res){
    var collected = this.data.collected 
    if(res.code == 2000){
      wx.showToast({
        title: collected? "取消成功" : "收藏成功",
        icon: 'success',
        duration: 1000
      })
      this.setData({
        collected:!collected,
      })
    } else if (res.code == 4003) {
      wx.showToast({
        title: '没有登录或登录过期',
        icon: "none",
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
  isCollectCallback: function (res) {
    if(res.code==2000){
      if(res.result==true){
        this.setData({
          collected:true
        })
      }
    }
  }
})