import requests
import json
from utils.config import *
ss = requests.session()
def getFoodCategory():
    data = {
    }
    url = URL.format(env=cur_url, term=client_url, action="/food/getFoodCategory")
    getReq(url,ss,data=data)

def searchFood():
    data = {
        # "keyword":"甜",
        "kind_id":0
    }
    url = URL.format(env=cur_url, term=client_url, action="/food/searchFood")
    getReq(url, ss, data=data)

def getFoodDetails():
    data = {
        "food_id": 2
    }
    url = URL.format(env=cur_url, term=client_url, action="/food/getFoodDetails")
    getReq(url, ss, data=data)


def commentFood():
    data = json.dumps({
        "token":"123",
        "food_id": 2,
        "content":"真香"
    })
    url = URL.format(env=cur_url, term=client_url, action="/food/commentFood")
    postReq(url, ss, data)

def getComment():
    data = {
        "food_id": 2
    }
    url = URL.format(env=cur_url, term=client_url, action="/food/getComment")
    getReq(url, ss, data=data)


def main():
    # getFoodCategory()
    # searchFood()
    # getFoodDetails()
    commentFood()
    getComment()


if __name__ == '__main__':
    main()