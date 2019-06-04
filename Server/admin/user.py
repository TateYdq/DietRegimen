import requests
import json
from utils.config import *
ss = requests.session()
def addUser():
    data = json.dumps({
        "token":"123",
        "name":"123",
        "age":60,
        "gender":"female",
        "user_image_path":"test",
        "diseases_focus":"123"
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addUser")
    postReq(url,ss,data)

def updateUser():
    data = json.dumps({
        "token":"123",
        "user_id":8,
        "name":"uuuu",
        "age":63,
        "gender":"male",
        "user_image_path":"haha",
        "diseases_focus":"高血脂"
    })
    url = URL.format(env=cur_url, term=admin_url, action="/updateUser")
    postReq(url,ss,data)


def main():
    addUser()

if __name__ == '__main__':
    main()