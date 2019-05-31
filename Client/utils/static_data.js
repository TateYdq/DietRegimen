var foodInfo=[
    {
    "foodID":1,
    "name":"土豆",
    "viewCount":"20",
    "foodKind":"蔬菜",
    "collectCount":"7",
    "info":"土豆，学名是马铃薯，拉丁名：Solanum tuberosum L\“土豆\”是通称，个别地区叫洋芋，在法国，土豆被称为“地下苹果”。",
    "photoPath":"../../../imgs/food/tudou.jpeg"
    },
    {
      "foodID": 2,
      "name": "苹果",
      "foodKind": "水果",
      "viewCount": "10",
      "collectCount": "3",
      "info":"苹果是一种低热量食物，每100克产生60千卡热量。苹果中营养成分可溶性大，易被人体吸收，故有“活水”之称。它有利于溶解硫元素，使皮肤润滑柔嫩 。",
      "photoPath": "../../../imgs/food/pingguo.jpeg"
    }
  ];
var foodKindInfo = [
  {
    "kindID":1,
    "kindName":"蔬菜",
    "kindInfo":"蔬菜是人们日常饮食中必不可少的食物之一。蔬菜可提供人体所必需的多种维生素和矿物质等营养物质。",
    "photoPath":"../../../imgs/food/shucai.jpg",
    "viewCount":3
  },
  {
    "kindID": 2,
    "kindName": "水果",
    "kindInfo": "水果是指多汁且主要味觉为甜味和酸味，可食用的植物果实。水果不但含有丰富的维生素营养，而且能够促进消化。",
    "photoPath": "../../../imgs/food/shuiguo.jpeg",
    "viewCount":"2"
  }
];


var diseaseInfo = [
{
    "diseaseID":1,
    "name":"高血压",
    "diseaseKind":"心血管内科",
    "viewCount":30,
    "collectCount":6,
    "info":"高血压是指以体循环动脉血压（收缩压和/或舒张压）增高为主要特征（收缩压≥140毫米汞柱，舒张压≥90毫米汞柱）",
    "photoPath": "../../../imgs/diseases/gaoxueya.jpg"
},
  {
    "diseaseID": 2,
    "name": "糖尿病",
    "diseaseKind": "心血管内科",
    "viewCount": 12,
    "collectCount": 3,
    "info": "糖尿病是一组以高血糖为特征的代谢性疾病。高血糖则是由于胰岛素分泌缺陷或其生物作用受损，或两者兼有引起。",
    "photoPath": "../../../imgs/diseases/tangniaobing.jpeg"
  }
];
var userInfo = [];
var foodCommentInfo = [];
var diseaseCommentInfo = [];
var userCollectDiseaseInfo = [];
var userCollectFoodInfo = [];
var diseaseFoodInfo =[
  {
    "id":1,
    "diseaseID":1,
    "diseaseName":"高血压",
    "foodID":3,
    "foodName":"豆腐",
  }
];
var questionInfo = [];


/*
暴露的接口
*/
module.exports = {
  foodInfo: foodInfo,
  foodKindInfo: foodKindInfo,
  diseaseInfo: diseaseInfo,
  userInfo: userInfo
}