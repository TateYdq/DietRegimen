import requests


staging_url = "http://127.0.0.1:8080"
online_url = "http://api.ydq6.com/DietRegimen"

admin_url="/control1/admin"

client_url="/client"


URL="{env}{term}{action}"

admin_token = "1wqe213"
user_token = "123"

cur_url = staging_url


def postReq(url,ss,data,token):
    print(url)
    headers = {
        'Content-Type': 'application/json; charset=UTF-8',
        "Token": token
    }
    res = ss.post(url, headers=headers, data=data)
    print(res.content.decode("utf-8"))


def getReq(url,ss,data,token=""):
    print(url)
    # headers = {'Content-Type': 'application/json; charset=UTF-8'}
    headers = {"Token": token}
    res = ss.get(url, params=data,headers=headers)
    print(res.content.decode("utf-8"))