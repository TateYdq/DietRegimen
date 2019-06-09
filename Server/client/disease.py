import requests
import json
from utils.config import *
ss = requests.session()
def getDiseasesLists():
    data = {
        # "keyword":""
        "keyword":"心脏"
    }
    url = URL.format(env=cur_url, term=client_url, action="/health/getDiseasesLists")
    getReq(url,ss,data=data)

def searchDisease():
    data = {
        # "keyword":"甜",
        "kind_id":0
    }
    url = URL.format(env=cur_url, term=client_url, action="/health/searchDisease")
    getReq(url, ss, data=data,token=user_token)

def getDiseaseDetails():
    data = {
        "disease_id": 2
    }
    url = URL.format(env=cur_url, term=client_url, action="/health/getDiseaseDetails")
    getReq(url, ss, data=data,token=user_token)


def commentDisease():
    data = json.dumps({
        "token":"123",
        "disease_id":3,
        "content":"难受"
    })
    url = URL.format(env=cur_url, term=client_url, action="/health/commentDisease")
    postReq(url, ss, data=data,token=user_token)

def getComment():
    data = {
        "disease_id": 4
    }
    url = URL.format(env=cur_url, term=client_url, action="/health/getComment")
    getReq(url, ss, data=data,token=user_token)


def main():

    # searchDisease()
    getDiseaseDetails()
    # getDiseasesLists()
    # commentDisease()
    # getComment()


if __name__ == '__main__':
    main()