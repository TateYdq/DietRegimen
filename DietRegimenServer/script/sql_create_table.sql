create database ydq6_miniprogram default character set utf8mb4;
use ydq6_miniprogram;

create table user_info(
user_id  int auto_increment primary key,
name varchar(50),
age int,
gender varchar(50),
user_image_path varchar(50),
diseases_focus text,
keywords  text    #用户关键词
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table food_info(
food_id int auto_increment primary key,
name varchar(50),
food_kind_id int,
food_kind  varchar(50),
info   text,       #食物介绍
effect  text,     #食物功效
keyword  text,         #食物关键词
taboo text,          #禁忌
view_count bigint,
collect_count bigint,
photo_path  varchar(50) ,      #食物图片路径，只有一张
voice_path varchar(50)       #语音路径
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table disease_info(
disease_id int auto_increment primary key,
name varchar(50),
disease_kind  varchar(50),
info text, #疾病介绍
photo_path  varchar(50) ,#疾病图片路径，只有一张
voice_path varchar(50),
view_count bigint,
collect_count bigint
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table food_comment_info(
id int auto_increment primary key,
food_id int,
user_id int,
user_name varchar(50),
comment text,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


create table disease_comment_info(
id int auto_increment primary key,
disease_id int,
user_id int,
user_name varchar(50),
comment text,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


create table food_kind_info(
kind_id int auto_increment primary key,
kind_name varchar(50),
kind_info text,
photo_path varchar(50),
view_count bigint
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


create table user_collect_food_info(
id int auto_increment primary key,
user_id int,
food_id int,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table user_collect_disease_info(
id int auto_increment primary key,
user_id int,
disease_id int,
record_time varchar(50)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table question_info(
id int auto_increment primary key,
info text,#问题内容
answer text,  #问题回复选项，一般只有四个选项A，B,C
response text  #问题选项对应的响应。
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table disease_food_info(
id int auto_increment primary key,
disease_id int,
disease_name varchar(50),
food_name varchar(50),
food_id int
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

