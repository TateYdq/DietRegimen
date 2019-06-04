import requests
import json
from utils.config import *
ss = requests.session()
def addDisease():
    data = json.dumps({
        "token":"123",
        "name":"糖尿病",
        "disease_kind":"心血管疾病",
        "info":"糖尿病很严重",
        "view_count":1,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addDisease")
    postReq(url,ss,data)

def updateDisease():
    data = json.dumps({
        "token":"123",
        "disease_id":3,
        "name":"心脏病",
        "food_kind":"水果",
        "info":"容易突发，需预防\n",
        "view_count":3,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/updateDisease")
    postReq(url,ss,data)


def main():
    addDisease()
    # updateDisease()
if __name__ == '__main__':
    main()