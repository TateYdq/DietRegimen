import requests


staging_url = "http://127.0.0.1:8080/DietRegimen"
online_url = "api.ydq6.com/DietRegimen"

admin_url="/control1/admin"

client_url="/client"


URL="{env}{term}{action}"


cur_url = staging_url


def postReq(url,ss,data):
    print(url)
    headers = {'Content-Type': 'application/json; charset=UTF-8'}
    res = ss.post(url, headers=headers, data=data)
    print(res.content.decode("utf-8"))


def getReq(url,ss,data):
    print(url)
    # headers = {'Content-Type': 'application/json; charset=UTF-8'}
    res = ss.get(url, params=data)
    print(res.content.decode("utf-8"))