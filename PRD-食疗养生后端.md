<center><h1>PRD-食疗养生后端</h1></center>

# 一、文档综述



## 1.版本修订记录

v0.1 创建文档



# 二、数据库

数据库类型:mysql

数据库版本号:待定

数据库库名:DietRegimen

## 2.1. 用户信息表

user_info(

​	userID  int auto_increment primary key,

​	name varchar(50),

​	age int,

   gender varchar(50),

​	userImagePath varchar(50),

   diseasesHistory text,

​	keywords  text    #用户关键词

​	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



## 2.2.食物信息表

food_info(

​	foodID int auto_increment primary key,

​    name varchar(50),

​	foodKindID int,

​	foodKind  varchar(50),     

​	info   text,       #食物介绍

​	effect  text,     #食物功效

​	keyword  text,         #食物关键词

​    taboo text,          #禁忌

​	viewCount bigint,	

​	collectCount bigint,

​	photoPath  varchar(50) ,      #食物图片路径，只有一张

​    voicePath varchar(50)       #语音路径

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.3 疾病信息表

disease_info(

​	diseaseID int auto_increment primary key,

​    name varchar(50),

​	diseaseKind  varchar(50),

​	info   text,       #疾病介绍

​	photoPath  varchar(50) ,      #疾病图片路径，只有一张

​	voicePath varchar(50),

​	viewCount bigint,

​    collectCount bigint

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



## 2.4 用户评论食物信息表

food_comment_info(

​	id int auto_increment primary key,

​    foodID int,

​	userID int,

​	userName varchar(50),

​	comment text,

​	date varchar(50),

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.5 用户评论疾病信息表

disease_comment_info(

​	id int auto_increment primary key,

​    diseaseID int,

​	userID int,

​	userName varchar(50),

​	comment text,

​	date varchar(50),

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.6 食物种类表

food_kind_info(

​		kindID int auto_increment primary key,

​		kindName varchar(50),

​		kindInfo   text,

​		photoPath varchar(50),

​		viewCount bigint,

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.7 用户收藏食物表

user_collect_food_info(

​    id int auto_increment primary key,

​	userID int,

​	foodID int,

​	date varchar(50)

))ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.8 用户收藏疾病表

user_collect_disease_info(

​    id int auto_increment primary key,

​	userID int,

​	diseaseID int,

​	date varchar(50)

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



## 2.9 答题信息表

question_info(

​	id int auto_increment primary key,

​	info text,      #问题内容

​	answer text,  #问题回复选项，一般只有四个选项A，B,C,D

​	response text  #问题选项对应的响应。具体内容TODO

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

## 2.10 疾病推荐食物表

disease_food_info{

​	id int auto_increment primary key,

​	diseaseID int,

​	diseaseName varchar(50),

​	foodName varchar(50),

​	foodID int,

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 三、URL接口

code 通用返回码

| Code | 说明       | 备注 |
| ---- | ---------- | ---- |
| 2000 | 成功       |      |
| 4003 | 禁止访问   |      |
| 5000 | 失败       |      |
| 5003 | 服务器异常 |      |
|      |            |      |





##3.1.管理员端接口

- 统一前缀:/DietRegimen/control1/admin/    #为了安全

- 为了简化，不设置权限分配，统一用token。(token不在文档中给出)

- 照片需要先上传得到路径后才增删修改，图片均保存为相对路径。
- 不用修改的字段默认不填。
- 所有请求必须给出token
- 修改，删除时必须给出对应id

### 3.1.1 用户信息表

#### 3.1.1.1 添加用户

POST /DietRegimen/control1/admin/addUser

Request Body:{

​	"token":"",    #此为管理员token

​	"name":"" ,

​     "age":0,     

​      "gender":"",

​	 "userImagePath":"",

​     "diseasesHistory":"",

}

Resonse Body:{

​	"code":""    #2000,4003,5000,5003

}

#### 3.1.1.2 修改用户信息

POST /DietRegimen/control1/admin/updateUser

Request Body:{

​	"token":"",    #此为管理员token

​	userID:0     #用户id

​	"name":"" ,

​     "age":0,     

​      "gender":"",

​	 "userImagePath":"",

​     "diseasesHistory":"",

​	  "keywords":"" 

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.1.3 删除用户信息

POST /DietRegimen/control1/admin/delUser

Request Body:{

​	"token":"",    #此为管理员token

​	userID:0    #用户id

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.2 食物信息表

#### 3.1.2.1 添加食物

POST /DietRegimen/control1/admin/addFood

Request Body:{

​	"token":"",    #此为管理员token

​	"name":"" ,

​	"foodKindID":0 ,

​	"foodKind":"" ,     

​	"info":"" ,

​	"effect":"",    

​	"keyword":"" ,  

​    "taboo":"",  

​	"viewCount":0,	

​	"photoPath":""  ,     

​    "voicePath":""     

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.2.1 修改食物信息

POST /DietRegimen/control1/admin/updateFood

Request Body:{

​	"token":"",    #此为管理员token

​	"foodID":0,   #食物id

​	"name":"" ,

​	"foodKindID":0 ,

​	"foodKind":"" ,     

​	"info":"" ,

​	"effect":"",    

​	"keyword":"" ,  

​    "taboo":"",  

​	"viewCount":0,	

​	"collectCount":0,

​	"photoPath":""  ,     

​    "voicePath":""     

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



#### 3.1.2.2 删除食物信息

POST /DietRegimen/control1/admin/updateFood

Request Body:{

​	"token":"",    #此为管理员token

​	"foodID":0,   #食物id

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.3 疾病信息表

#### 3.1.3.1 添加疾病

POST /DietRegimen/control1/admin/addDisease

Request Body:{

​	"token":"",    #此为管理员token

​    "name":"",

​	"diseaseKind":"",

​	"voicePath":"",

​	"viewCount":0 ,

​    ”collectCount“:0 ,

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.3.2 修改疾病

POST /DietRegimen/control1/admin/updateDisease

Request Body:{

​	"token":"",    #此为管理员token

​	"diseaseID":0,   #疾病id

​    "name":"",

​	"diseaseKind":"",

​	"voicePath":"",

​	"viewCount":0 ,

​    ”collectCount“:0 ,

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.3.2 删除疾病

POST /DietRegimen/control1/admin/updateDisease

Request Body:{

​	"token":"",    #此为管理员token

​	"diseaseID":0,   #疾病id

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

### 3.1.4 用户评论食物信息表(暂不提供)

#### 3.1.4.1 添加用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.4.2 修改用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.4.3 删除用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

### 3.1.5 用户评论疾病信息表(暂不提供)

#### 3.1.5.1 添加

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.5.2 修改

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.5.3 删除

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.6 食物种类表

#### 3.1.6.1 添加

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.6.2 修改

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"kindID":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.6.3 删除

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"kindID":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.7 用户收藏食物表(暂不提供)

#### 3.1.7.1 添加

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.7.2 修改

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.7.3 删除

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.8 用户收藏疾病表(暂不提供)

#### 3.1.8.1 添加

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.8.2 修改

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.8.3 删除

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.9 答题信息表

#### 3.1.9.1 添加

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.9.2 修改

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.9.3 删除

POST /DietRegimen/control1/admin/

Request Body:{

​	"token":"",    #此为管理员token

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.10 图片相关

#### 3.1.10.1  上传图片

POST /DietRegimen/control1/admin/uploadImage

Request Body:{

​	"token":"",    #此为管理员token

​	"image":""

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

​	 "path":""            #图片路径

}

#### 3.1.10.1  下载图片

POST /DietRegimen/control1/admin/downloadImage

Request Body:{

​	"token":"",    #此为管理员token

​	"path":""       #图片相对路径

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

 	"image":""        #图片  

}

#### 3.1.10.1  修改图片

POST /DietRegimen/control1/admin/updateImage

Request Body:{

​	"token":"",    #此为管理员token

​	 "path":""            #图片路径

​	"image":""

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

## 3.2.普通用户接口(给小程序提供的接口)

统一前缀:/DietRegimen/client

### 3.2.1 用户相关

#### 3.2.1.1 获取个人信息

POST  /DietRegimen/client/user/getUsertInfo

Request Body:{

​	token                      用户token

}

Resonse Body:{

​		code              返回码

​		userInfo:{}         用户详细信息

}

#### 3.2.1.2 获取收藏信息

POST  /DietRegimen/client/user/getCollectInfo

Request Body:{

​	token                      用户token

}

Resonse Body:{

​	code                  返回码

​    collectInfo:{}      收藏信息，包含疾病和食物信息

}

#### 3.2.1.3 个人信息设置

POST /DietRegimen/client/user/postUsertInfo

Request Body:{

​	token                      用户token

​	userInfo                  用户信息

}

Resonse Body:{

body:{

​		code              返回码

}

#### 3.2.1.4 登录

POST /DietRegimen/client/user/userLogin

Request Body:{

​	code                      登录凭证

​	userInfo                  AppID

​	secret             

​	grant_type

}

Resonse Body:{

body:{

​		code              返回码

​		sessionID 

}

[注]此处可参考小程序的登录API

### 3.2.2 食物相关

####  3.2.2.1 获取食物详细信息

POST /DietRegimen/client/food/getFoodDetails

Request Body:{

​		foodID           食物id

}

Resonse Body:{

​		code:                  返回代码

​		details:                 食物详细信息

}



#### 3.2.2.2 获取食物种类

GET /DietRegimen/client/food/getFoodCategory

Resonse Body:{

​	code:               返回代码

​	foodCategory :{}  食物种类列表

}

#### 3.2.2.3 搜索食物清单/获取指定食物种类的食物清单

POST /DietRegimen/client/food/searchFood

Request Body:{

​		keywords:              关键词

}

Resonse Body:{

​		code:                    返回代码

​		foods:{}                食物集合

}



#### 3.2.2.4 发表评论

POST  /DietRegimen/client/food/commentFood

Request Body:{

​	token                 用户token

​	foodID               食物id

   content               评论内容

}

Resonse Body:{

​		code         返回代码

}



#### 3.2.2.5 获取语音播报

POST  /DietRegimen/client/food/foodVoice

Request Body:{

​	token         用户token

​	foodID       食物id

}

Resonse Body:{

​	foodVoice.mp3     语音文件

}

通过http返回状态码判获取断成功与失败。



#### 3.2.2.6 收藏食物

POST  /DietRegimen/client/user/collectFood

Request Body:{

​	token              用户token

   foodID             食物id

}



Resonse Body:{

​	code                  返回码

}

### 3.2.3 疾病相关

#### 3.2.3.1 获取疾病详细信息

POST /DietRegimen/client/health/getDiseaseDetails

Request Body:{

​	diseaseID  疾病id

}

Resonse Body:{

​	code   返回代码

​	diseaseDetails     疾病详细信息

}



#### 3.2.3.2 获取疾病清单

GET /DietRegimen/client/health/getDiseasesLists

Resonse Body:{

​		code          返回代码

​		diseases:{}      疾病清单

}

#### 3.2.3.3 发表评论

POST  /DietRegimen/client/health/commentDisease

Request Body:{

​	token                 用户token

​	diseaseID           食物id

​    content              评论内容

}

Resonse Body:{

body:{

​		code         返回代码

}

#### 3.2.3.4 获取语音播报

POST  /DietRegimen/client/health/diseaseVoice

Request Body:{

​	token             用户token

​	diseaseID       疾病id

}

Resonse Body:{

diseaseVoice.mp3     语音文件

}

通过http返回状态码判获取断成功与失败。

#### 3.2.3.5 收藏疾病

POST  /DietRegimen/client/user/collecDisease

Request Body:{

​	token                  用户token

   diseaseID            疾病id

}

Resonse Body:{

​	code                  返回码

}

### 3.2.3 推荐相关

#### 3.2.3.1 获取推荐信息

POST /DietRegimen/client/recommend/getRecInfo

Request Body:{

​			token           用户token

  }

Resonse Body:{

​		code               返回码

​		info:{}             推荐信息

​		foods:{}          食物清单

}




### 3.2.4 问卷相关

#### 3.2.4.1 获取问卷信息

GET  /DietRegimen/client/recommend/getQuestionnaireInfo

Resonse Body:{

​	code          返回码

​    infoLists     问卷列表     

}



#### 3.2.4.2 提交问卷

POST  /DietRegimen/client/recommend/submitQuestionnaire

Request Body:{

​	token                      用户token

​	answerInfoLists      问卷回答信息

}

Resonse Body:{

​		code       返回码

}

