import requests
import json
from utils.config import *
ss = requests.session()
def addFood():
    data = json.dumps({
        "name":"养生菇粥",
        # "food_kind_id":1,
        # "food_kind":"水果",
        "info":"做法：1.羊肉洗净，切块；虾子，洗净；山药洗净，切片备用。",
        "effect":"安定凝神、利尿通便、健胃消食。",
        "view_count":0,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addFood")
    postReq(url,ss,data,admin_token)

def updateFood():
    data = json.dumps({
        "food_id":2,
        "name": "西瓜",
        "Keyword":"好吃",
        "view_count": 15,
    })
    url = URL.format(env=cur_url, term=admin_url, action="/updateFood")
    postReq(url,ss,data,admin_token)


def main():
    addFood()
    # updateFood()
if __name__ == '__main__':
    main()