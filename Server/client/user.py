import requests
import json
from utils.config import *
ss = requests.session()
def getUserInfo():
    data = {
        "token":"13"
    }
    url = URL.format(env=cur_url, term=client_url, action="/user/getUserInfo")
    getReq(url, ss, data=data,token=user_token)


def updateUserInfo():
    data = json.dumps({
        "token":"123",
        "user_info":{
            "user_id": 3,
            "name": "uuuu",
            "age": 63,
            "gender": "male",
            "user_image_path": "haha",
            "diseases_focus": "高血脂"
        }
    })
    url = URL.format(env=cur_url, term=client_url, action="/user/updateUserInfo")
    postReq(url,ss,data)


def collectFood():
    data = json.dumps({
        "token":"123",
        "food_id":1
    })
    url = URL.format(env=cur_url, term=client_url, action="/user/collectFood")
    postReq(url, ss, data,token=user_token)

def getCollectFood():
    data = {
        "token":"13"
    }
    url = URL.format(env=cur_url, term=client_url, action="/user/getCollectFood")
    getReq(url, ss, data=data,token=user_token)

def collectDisease():
    data = json.dumps({
        "token":"123",
        "disease_id":1
    })
    url = URL.format(env=cur_url, term=client_url, action="/user/collectDisease")
    postReq(url, ss, data,token=user_token)

def getCollectDisease():
    data = {
        "token":"13"
    }
    url = URL.format(env=cur_url, term=client_url, action="/user/getCollectDisease")
    getReq(url, ss, data=data,token=user_token)



def login():
    data = json.dumps({
        "code":"081oGCV92py1zK0obwW92gBsV92oGCVQ",
    })
    url = URL.format(env=cur_url, term=client_url, action="/user/userLogin")
    postReq(url, ss, data,token=user_token)




def main():
    login()
    # getUserInfo()
    # collectFood()
    # getCollectFood()
    # collectDisease()
    # getCollectDisease()
if __name__ == '__main__':
    main()