// pages/test/test.js
// const model = require('../cityChoose/cityChoose.js')
// const config = require('../../utils/config.js')
// const util = require('../../utils/util.js')
// const app = getApp()
// var mydata = {
//   end: 0,
//   replyUserName: ""
// }
Page({

  /**
   * 页面的初始数据
   */
  data: {
    list: [
      {
        id: 1,
        userPhoto: '../../imgs/images/tuiList1.jpg',
        userName: '曹总',
        replyUserName: ""
        // collectCount: 65,
        // kindInfo: '功效：安定凝神、利尿通便、健胃消食。'
      },
      {
        id: 2,
        userPhoto: '../../imgs/images/tuiList2.jpg',
        userName: 'CFO',
        replyUserName: ''
        // collectCount: 69,
        // kindInfo: '功效：养胃生津、除烦解渴、利尿通便、清热解毒。'
      }, {
        id: 3,
        userPhoto: '../../imgs/images/tuiList3.jpg',
        userName: '八宝养生饭',
        replyUserName: ''
        // collectCount: 15,
        // kindInfo: '功效：减肥排毒、润肺止咳、安定凝神。'
      }
    ], 
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    // var that = this;
    // mydata.sourceId = options.sourceId
    // mydata.commentId = "";
    // mydata.replyUserName = "";
    //设置scroll的高度
    wx.getSystemInfo({
      success: function (res) {
        that.setData({
          scrollHeight: res.windowHeight,
          userId: app.globalData.haulUserInfo.id
        });
      }
    });
    // mydata.page = 1;
    // that.getPageInfo(mydata.page);
  },
//   /**
//    * 页面下拉刷新事件的处理函数
//    */
//   refresh: function () {
//     console.log('refresh');
//     mydata.page = 1
//     this.getPageInfo(mydata.page, function () {
//       this.setData({
//         list: []
//       })
//     });
//     mydata.end = 0;
//   },
//   /**
//    * 页面上拉触底事件的处理函数
//    */
//   bindDownLoad: function () {
//     console.log("onReachBottom");
//     var that = this;
//     if (mydata.end == 0) {
//       mydata.page++;
//       that.getPageInfo(mydata.page);
//     }
//   },
//   bindReply: function (e) {
//     console.log(e);
//     mydata.commentId = e.target.dataset.commentid;
//     mydata.replyUserName = e.target.dataset.commentusername;
//     this.setData({
//       replyUserName: mydata.replyUserName,
//       reply: true
//     })
//   },
//   // 合并数组
//   addArr(arr1, arr2) {
//     for (var i = 0; i < arr2.length; i++) {
//       arr1.push(arr2[i]);
//     }
//     return arr1;
//   },
//   deleteComment: function (e) {
//     console.log(e);
//     var that = this;
//     var commentId = e.target.dataset.commentid;

//     wx.showModal({
//       title: '删除评论',
//       content: '请确认是否删除该评论？',
//       success: function (res) {
//         if (res.confirm) {
//           wx.request({
//             url: config.deleteComment,
//             method: "POST",
//             data: {
//               commentId: commentId
//             },
//             header: {
//               "content-type": "application/x-www-form-urlencoded;charset=utf-8",
//             },
//             success: res => {
//               that.refresh();
//               wx.showToast({
//                 title: "删除成功"
//               })
//             }
//           })
//         } else if (res.cancel) {
//           console.log('用户点击取消')
//         }
//       }
//     })
//   },
//   cancleReply: function (e) {
//     mydata.commentId = "";
//     mydata.replyUserName = "";
//     this.setData({
//       replyUserName: mydata.replyUserName,
//       reply: false
//     })
//   },
//   // 更新页面信息
//   // 此处的回调函数在 传入新值之前执行 主要用来清除页面信息
//   getPageInfo(page, callback) {
//     var that = this;
//     util.showLoading();
//     console.log("getPageInfo");
//     console.log("page" + page);
//     var limited = 6;
//     var offset = (page - 1) * 6;
//     wx.request({
//       url: config.getComments,
//       method: "POST",
//       data: {
//         sourceId: mydata.sourceId,
//         limited: limited,
//         offset: offset
//       },
//       header: {
//         "content-type": "application/x-www-form-urlencoded;charset=utf-8",
//       },
//       success: res => {
//         console.log(res);
//         if (page == 1) {
//           that.data.list = res.data;
//           that.setData({
//             list: that.data.list
//           })
//           mydata.end = 0;
//         } else {
//           // 当前页为其他页
//           var list = that.data.list;
//           if (res.data.length != 0) {
//             list = that.addArr(list, res.data);
//             that.setData({
//               list: list
//             })
//             mydata.end = 0;
//           } else {
//             mydata.end = 1;
//           }
//         }
//         wx.hideLoading();
//       }
//     })
//   },
//   submitForm(e) {
//     var form = e.detail.value;
//     var that = this;
//     console.log(app.globalData.haulUserInfo);
//     if (form.comment == "") {
//       util.showLog('请输入评论');
//       return;
//     }
//     // 提交评论
//     wx.request({
//       url: config.insertComment,
//       method: "POST",
//       data: {
//         sourceId: mydata.sourceId,
//         comment: form.comment,
//         userId: app.globalData.haulUserInfo.id,
//         userName: app.globalData.haulUserInfo.userName,
//         replyCommentId: mydata.commentId,
//         replyUserName: mydata.replyUserName,
//         userPhoto: app.globalData.haulUserInfo.userPhoto
//       },
//       header: {
//         "content-type": "application/x-www-form-urlencoded;charset=utf-8",
//         //token: app.globalData.token
//       },
//       success: res => {
//         console.log(res)
//         if (res.data.success) {
//           wx.showToast({
//             title: "回复成功"
//           })
//           that.refresh();
//           mydata.commentId = "";
//           mydata.replyUserName = "";
//           this.setData({
//             replyUserName: mydata.replyUserName,
//             reply: false
//           })
//         } else {
//           wx.showToast({
//             title: '回复失败，请检查您的网络',
//           })
//         }
//       }
//     })
//   }
})



// Page({

//   /**
//    * 页面的初始数据
//    */
//   data: {

//   },

//   /**
//    * 生命周期函数--监听页面加载
//    */
//   onLoad: function (options) {

//   },

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

//   /**
//    * 用户点击右上角分享
//    */
//   onShareAppMessage: function () {

//   }
// })