import requests
import json
from utils.config import *
ss = requests.session()
def addFoodKind():
    data = json.dumps({
        "token":"123",
        "kind_name":"水果",
        "kind_info":"水果很好吃",
        "view_count":3,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addFoodKind")
    postReq(url,ss,data)

def updateFoodKind():
    data = json.dumps({
        "token":"123",
        "kind_id":1,
        "kind_name":"水果",
        "kind_info":"水果当然很好吃",
        "view_count":6,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/updateFoodKind")
    postReq(url,ss,data)


def main():
    addFoodKind()
    updateFoodKind()
if __name__ == '__main__':
    main()