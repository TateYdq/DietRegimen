<center><h1>PRD-食疗养生后端</h1></center>

# 一、文档综述



## 1.版本修订记录

v0.1 创建文档

v0.2 由于后端采用gin所以所有命名方式由小驼峰换为全小写模式。

v0.3 修改疾病史字段，改为关注疾病

v0.4 决定所有Get请求默认为参数形式，所有Post请求默认为json格式。将taboo字段从食物信息表中去掉，并加到疾病信息表中

V0.5 小程序和后端交互决定采用token方式，并且把token放在请求头中传输

v0.6 修改文档中的错误的地方和增加了2.11，2.12，2.13表



# 二、数据库

数据库类型:mysql

数据库版本号:待定

数据库库名:DietRegimen

## 2.1. 用户信息表

 ```
create table user_info(
user_id  int auto_increment primary key,
open_id varchar(50),
name varchar(50),
age int,
gender varchar(50),
user_image_path varchar(50),
diseases_focus text,
keywords  text,   #用户关键词
update_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
 ```





## 2.2.食物信息表

```
create table food_info(
food_id int auto_increment primary key,
name varchar(50),
food_kind_id int,
food_kind  varchar(50),
info   text,       #食物介绍
effect  text,     #食物功效
keyword  text,         #食物关键词
view_count bigint,
collect_count bigint,
photo_path  varchar(50) ,      #食物图片路径，只有一张
voice_path varchar(50)       #语音路径
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.3 疾病信息表

```
create table disease_info(
disease_id int auto_increment primary key,
name varchar(50),
disease_kind  varchar(50),
info text, #疾病介绍
taboo text,  #禁忌
photo_path  varchar(50) ,#疾病图片路径，只有一张
voice_path varchar(50),
view_count bigint,
collect_count bigint
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```





## 2.4 用户评论食物信息表

```
create table food_comment_info(
id int auto_increment primary key,
food_id int,
user_id int,
user_name varchar(50),
comment text,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.5 用户评论疾病信息表

```
create table disease_comment_info(
id int auto_increment primary key,
disease_id int,
user_id int,
user_name varchar(50),
comment text,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.6 食物种类表

```
create table food_kind_info(
kind_id int auto_increment primary key,
kind_name varchar(50),
kind_info text,
photo_path varchar(50),
view_count bigint
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.7 用户收藏食物表

```
create table user_collect_food_info(
id int auto_increment primary key,
user_id int,
food_id int,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.8 用户收藏疾病表

```
create table user_collect_disease_info(
id int auto_increment primary key,
user_id int,
disease_id int,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.9 问题表

```
create table question_info(
question_id int auto_increment primary key,
info text,      #问题内容
answer_a text,  #问题A
response_a text,  #问题A对应的响应。
answer_b text,  #问题A
response_b text,  #问题A对应的响应。
answer_c text,  #问题A
response_c text  #问题A对应的响应。
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```





## 2.10 答题表

```
create table answer_sheet(
id int auto_increment primary key,
question_id int,   #问题id
user_id int,       #用户id
answer text,  #答案
record_time varchar(50)        #记录时间
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 2.11 疾病推荐食物表

```
create table disease_food_rec(
id int auto_increment primary key,
disease_id int,
food_id int,
food_name varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 2.12 用户食物关系表

```
create table user_food_relation(
user_id int,
food_id int,
food_name varchar(50),
score int #相关度分值,推荐的时候按此排序
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 2.13 用户疾病关系表

```
create table user_disease_relation(
user_id int,
disease_id int,
disease_name varchar(50),
score int #相关度分值,推荐的时候按此排序
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```





# 三、URL接口

code 通用返回码

| Code | 说明       | 备注 |
| ---- | ---------- | ---- |
| 2000 | 成功       |      |
| 4003 | 禁止访问   |      |
| 5000 | 失败       |      |
| 5003 | 服务器异常 |      |
|      |            |      |



所有Get请求默认为参数形式，所有Post请求默认为json格式。

##3.1.管理员端接口

- 统一前缀:/DietRegimen/control1/admin/    #为了安全

- 为了简化，不设置权限分配，统一用token。(token不在文档中给出)

- 照片需要先上传得到路径后才增删修改，图片均保存为相对路径。
- 不用修改的字段默认不填。
- 所有请求Token均在header
- 修改，删除时必须给出对应id

### 3.1.1 用户信息表

#### 3.1.1.1 添加用户

- [x] POST /DietRegimen/control1/admin/addUser

Request Body:{

​	"name":"" ,

​     "age":0,     

​      "gender":"male",       #male或female

​	 "user_image_path":"",

​     "diseases_focus":"",

​	"keywords":"",

}

Resonse Body:{

​	"code":    

​	"user_id": 

}

#### 3.1.1.2 修改用户信息

- [x] POST /DietRegimen/control1/admin/updateUser

Request Body:{

​	user_id:0     #用户id

​	"name":"" ,

​     "age":0,     

​      "gender":"",

​	 "user_image_path":"",

​     "diseases_focus":"",

​	  "keywords":"" 

}

Resonse Body:{

​	"code":    

}

#### 3.1.1.3 删除用户信息[暂不提供]

POST /DietRegimen/control1/admin/delUser

Request Body:{

​	user_id:0    #用户id

}

Resonse Body:{

​	"code":

}



### 3.1.2 食物信息表

#### 3.1.2.1 添加食物

- [x] POST /DietRegimen/control1/admin/addFood

Request Body:{

​	"name":"" ,

​	"food_kind_id":0 ,

​	"foodKind":"" ,     

​	"info":"" ,

​	"effect":"",    

​	"keyword":"" ,  

​    "taboo":"",  

​	"viewCount":0,	

​	"photo_path":""  ,     

​    "voicePath":""     

}

Resonse Body:{

​	"code": 

​	"food_id":

}

#### 3.1.2.1 修改食物信息

- [x] POST /DietRegimen/control1/admin/updateFood

Request Body:{

​	"food_id":0,   #食物id

​	"name":"" ,

​	"food_kind_id":0 ,

​	"foodKind":"" ,     

​	"info":"" ,

​	"effect":"",    

​	"keyword":"" ,  

​    "taboo":"",  

​	"viewCount":0,	

​	"collect_count":0,

​	"photo_path":""  ,     

​    "voicePath":""     

}

Resonse Body:{

​	"code":

}



#### 3.1.2.2 删除食物信息[暂不提供]

POST /DietRegimen/control1/admin/uprecordTimeFood

Request Body:{

​	"food_id":,   #食物id

}

Resonse Body:{

​	"code":

}



### 3.1.3 疾病信息表

#### 3.1.3.1 添加疾病

- [x] POST /DietRegimen/control1/admin/addDisease

Request Body:{

​    "name":"",

​	"disease_kind":"",

​	"voicePath":"",

​	"viewCount":0 ,

​    ”collect_count“:0 ,

}

Resonse Body:{

​	"code":

}

#### 3.1.3.2 修改疾病

- [x] POST /DietRegimen/control1/admin/updateDisease

Request Body:{

​	"disease_id":0,   #疾病id

​    "name":"",

​	"disease_kind":"",

​	"voicePath":"",

​	"viewCount":0 ,

​    ”collect_count“:0 ,

}

Resonse Body:{

​	"code":

}

#### 3.1.3.3 添加疾病推荐食物

- [x] POST /DietRegimen/control1/admin/addDiseaseFoodRec

Request Body:{

​	"disease_id":,       #疾病id

​	"food_names":[]  #食物名字列表

}

Resonse Body:{

​	"code":

}



#### 3.1.3.4 删除疾病[暂不提供]

POST /DietRegimen/control1/admin/deleteDisease

Request Body:{

​	"disease_id":,   #疾病id

}

Resonse Body:{

​	"code":

}

### 3.1.4 用户评论食物信息表(暂不提供)

#### 3.1.4.1 添加用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.4.2 修改用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

​	"id":

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.4.3 删除用户评论食物

POST /DietRegimen/control1/admin/

Request Body:{

​	"id":

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

- [x] POST /DietRegimen/control1/admin/addFoodKind

Request Body:{

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.6.2 修改

- [x] POST /DietRegimen/control1/admin/updateFoodKind

Request Body:{

​	"kindID":

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.6.3 删除[暂不提供]

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

- [x] POST /DietRegimen/control1/admin/addQuestion

Request Body:{
  "info":""
  "answer_a":""'
  "response_a":"" 
  "answer_b":""
  "response_b":"" 
  "answer_c":"" 
  "response_c":""

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

​	"question_id":""  

}

#### 3.1.9.2 修改(暂不提供)

POST /DietRegimen/control1/admin/updateQuestion

Request Body:{

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

#### 3.1.9.3 删除(暂不提供)

POST /DietRegimen/control1/admin/deleteQuestion

Request Body:{

​	"id":0

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}



### 3.1.10 图片相关

#### 3.1.10.1  上传图片

POST /DietRegimen/control1/admin/uploadImage

Request Header:{

​	"token":"",    #此为管理员token

}

Request Body:{

​	file:              #文件

​	"num"          #编号,1为用户，2为食物，3为食物种类，4为疾病

​	"id"              #id号对应num,可能为user_id，food_id,kind_id,disease_id,

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

​	 "path":""            #图片路径

}

eg.curl -X POST 127.0.0.1:8080/file/fileUpload -F "file=@/Users/yudaoqing/Pictures/pingguo.jpeg" -F "num=3" -F "id=3" -H "Content-Type: multipart/form-data"



curl -X POST https://api.ydq6.com/DietRegimen//control1/admin/uploadImage -F "file=@/Users/yudaoqing/go/src/github.com/TateYdq/DietRegimen/Client/imgs/diseases/tangniaobing.jpeg" -F "num=4" -F "id=1" -H "Content-Type: multipart/form-data"

#### 3.1.10.1  下载图片

POST /DietRegimen/control1/admin/downloadImage

Request Body:{

​	"path":""       #图片相对路径

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

 	"image":""        #图片  

}

#### 3.1.10.1  修改图片

POST /DietRegimen/control1/admin/uprecord_timeImage

Request Body:{

​	 "path":""            #图片路径

​	"image":""

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

}

## 3.2.普通用户接口(给小程序提供的接口)

统一前缀:/DietRegimen/client

所有请求Token均在Rquest Header

所有Get请求默认为参数形式，所有Post请求默认为json格式。

### 3.2.1 用户相关

#### 3.2.1.1 登录

- [x] POST /DietRegimen/client/user/userLogin

Request Body:{

​	code                      登录凭证  

}

Resonse Body:{

body:{

​		code              返回码  

​		token  

​		session_id 

}

[注]此处可参考小程序的登录API

目前问题，照片问题。

#### 3.2.1.2 获取个人信息

- [x] POST  /DietRegimen/client/user/getUsertInfo

Request Body:{

​	token                      用户token

}

Resonse Body:{

​		code              返回码

​		user_info:{}         用户详细信息

}





#### 3.2.1.3 获取收藏食物信息

- [x] GET   /DietRegimen/client/user/getCollectFood

Request Body:{

​	token                      用户token

}

Resonse Body:{

​	code                  返回码

​    collect_info:{}      收藏食物信息

}

#### 3.2.1.4 获取收藏疾病信息

- [x] GET  /DietRegimen/client/user/getCollectDisease

Request Body:{

​	token                      用户token

}

Resonse Body:{

​	code                  返回码

​    collect_info:{}      收藏疾病信息

}

#### 3.2.1.5 个人信息设置

- [x] POST /DietRegimen/client/user/updateUserInfo

Request Body:{

​	token                      用户token

​	user_info                  用户信息

}

Resonse Body:{

body:{

​		code              返回码

}



#### 3.2.1.6 收藏食物

POST  /DietRegimen/client/user/collectFood

Request Body:{

​	token              用户token

   food_id             食物id

}



Resonse Body:{

​	code                  返回码

}



#### 3.2.1.7 收藏疾病

- [x] POST  /DietRegimen/client/user/collecDisease

Request Body:{

   disease_id            疾病id

}

Resonse Body:{

​	code                  返回码

}



#### 3.2.1.8 上传图片

- [x] POST  /DietRegimen/client/user/uploadUserImage

Request Body:{

​	file:              #文件

}

Resonse Body:{

​	"code":"2000"    #2000,4003,5000,5003

​	 "path":""            #图片路径

}

eg.curl -X POST 127.0.0.1:8080/client/user/uploadUserImage -F "file=@/Users/yudaoqing/Pictures/pingguo.jpeg" -H "token:123"  -H "Content-Type: multipart/form-data"

[注]注意登录时候需要上传用户图片。

### 3.2.2 食物相关

####  3.2.2.1 获取食物详细信息

- [x] GET /DietRegimen/client/food/getFoodDetails

Request Body:{

​		food_id         食物id

}

Resonse Body:{

​		code:                  返回代码

​		food_detail:                 食物详细信息

}



#### 3.2.2.2 获取食物种类

- [x] GET /DietRegimen/client/food/getFoodCategory

Request Body:{

​	无

}

Resonse Body:{

​	code:               返回代码

​	food_category :{}  食物种类列表

}

#### 3.2.2.3 搜索食物清单/获取指定食物种类的食物清单

- [x] Get /DietRegimen/client/food/searchFood

Request Body:{

​		keyword:              关键词(单关键词，可不填写)

​		kindID:                  种类ID(可不填写)

}

Resonse Body:{

​		code:                    返回代码

​		food_list:{}                食物集合

}



#### 3.2.2.4 发表评论

- [x] POST  /DietRegimen/client/food/commentFood

Request Body:{

​	token                 用户token

​	food_id               食物id

   content               评论内容

}

Resonse Body:{

​		code         返回代码

}

#### 3.2.2.5 获取评论

- [x] POST  /DietRegimen/client/food/getComment

Request Body:{

​	food_id               食物id

}

Resonse Body:{

​		code         返回代码

​	   comment_list:{}      评论列表

}

#### 3.2.2.6 获取语音播报

用下载音频接口





### 3.2.3 疾病相关

#### 3.2.3.1 获取疾病详细信息

- [x] GET /DietRegimen/client/health/getDiseaseDetails

Request Body:{

​	disease_id  疾病id

}

Resonse Body:{

​	code   返回代码

​	disease_detail     疾病详细信息

}



#### 3.2.3.2 获取疾病清单

- [x] GET /DietRegimen/client/health/getDiseasesLists

Request Body:{

​	keyword      可以不填，不填默认为空，即搜索所有疾病

}

Resonse Body:{

​		code          返回代码

​		disease_list:{}      疾病清单

}

#### 3.2.3.3 发表评论

- [x] POST  /DietRegimen/client/health/commentDisease

Request Body:{

​	token                 用户token

​	disease_id          食物id

​    content              评论内容

}

Resonse Body:{

body:{

​		code         返回代码

}



#### 3.2.3.5 获取评论

- [x] POST  /DietRegimen/client/health/getComment

Request Body:{

​	disease_id          疾病id

}

Resonse Body:{

​		code         返回代码

​	   comment_list:{}      评论列表

}

#### 3.2.3.6 获取语音播报

用下载音频接口



### 3.2.3 推荐相关

#### 3.2.3.1 获取推荐信息

- [x] POST /DietRegimen/client/recommend/getRecInfo

Request Body:{

​			token           用户token

  }

Resonse Body:{

​		code               返回码

​		food_list:{}             食物清单

​		disease_list:{}          推荐疾病清单

}




### 3.2.4 问卷相关

#### 3.2.4.1 获取问卷信息

- [x] GET  /DietRegimen/client/recommend/getQuestionnaire

Resonse Body:{

​	code          返回码

​    question_lists:[]     问卷列表     

}

[注]此请求必须带token

问卷列表数组每个元素包括,info,answer_a,response_a,answer_b,response_b,answer_c,response_c




#### 3.2.4.2 提交问卷

- [x] POST  /DietRegimen/client/recommend/submitQuestionnaire

Request Body:{

​	answer_sheets:[]      问卷回答信息

}

Resonse Body:{

​		code       返回码

}



answer_sheets为数组，每个元素包括question_id,answer,record_time。

其中answer为选项对应的响应。



### 3.2.5 下载图片

- [x] GET /DietRegimen/file/getImage

Request Body:{

​	"path":""       #图片路径

}

返回 ：文件



[注]用户的头像用这个接口下载

### 3.2.6 下载音频

- [x] GET /DietRegimen/file/getVoice

Request Body:{

​	"path":""       #图片路径,path为数据库中存储的path

}

返回 ：文件



# 四、其他问题

## 4.1 session,token,user_id问题

基于安全性考虑，使用token，不适用session,JWT。

随机生成key, 关联openid，存入redis中，当请求带入key，直接从redis中获取openid得到当前用户信息，这个其实也就是我们自己去维护了会话信息



具体机制:

客户端利用微信自带wx.login接口获取code

将code传给服务器端

服务器端根据openID查找用户，

若无则根据openID创建用户并得到userID，如果用户已经存在直接得到用户的userID.

生成token,将token-userID存储在缓存里，缓存过期时间为24小时。

最后返回token

从客户端获取token加密保存在本地缓存里，以后每次会话都放在header里。





## 4.2 图片问题



待调研决定