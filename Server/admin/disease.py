import requests
import json
from utils.config import *
ss = requests.session()
def addDisease():
    data = json.dumps({
        "name":"风湿病",
        "disease_kind":"风湿免疫内科",
        "info":"风湿病是一组侵犯关节、骨骼、肌肉、血管及有关软组织或结缔组织为主的疾病，其中多数为自身免疫性疾病。发病多较隐蔽而缓慢，病程较长，且大多具有遗传倾向。以关节痛、畏风寒为主症的一组极其常见的临床症候群。\n风湿病是风湿性疾病的简称，泛指影响骨、关节、肌肉及其周围软组织，如滑囊、肌腱、筋膜、血管、神经等一大组疾病。",
        "view_count":0,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addDisease")
    postReq(url,ss,data,admin_token)

def updateDisease():
    data = json.dumps({
        "token":"1wqe213",
        "disease_id":3,
        "name":"心脏病",
        "food_kind":"水果",
        "info":"容易突发，需预防\n",
        "view_count":3,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/updateDisease")
    postReq(url,ss,data,admin_token)

def addDiseaseFoodRec():
    data = json.dumps({
        "disease_id":1,
        "food_names":["西瓜","菠萝"],
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addDiseaseFoodRec")
    postReq(url,ss,data,admin_token)

def main():
    addDisease()
    # updateDisease()
    # addDiseaseFoodRec()
if __name__ == '__main__':
    main()