#POST /DietRegimen/control1/admin/uploadImage

#Request Header:{

​#	"token":"",    #此为管理员token

#}

#Request Body:{

​#	file:              #文件

​#	"num"          #编号,1为用户，2为食物，3为食物种类，4为疾病

#​	"id"              #id号对应num,可能为user_id，food_id,kind_id,disease_id,

#}

#Resonse Body:{

​#	"code":"2000"    #2000,4003,5000,5003

​#	 "path":""            #图片路径

#}






#curl -X POST https://api.ydq6.com/DietRegimen//control1/admin/uploadImage -F "file=@/Users/yudaoqing/go/src/github.com/TateYdq/DietRegimen/Client/imgs/images/tuiList1.jpg" -F "num=2" -F "id=4" -H "Content-Type: multipart/form-data" -H "Token: 1wqe213"


