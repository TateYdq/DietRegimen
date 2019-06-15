import requests
import json
from utils.config import *
ss = requests.session()
def addDisease():
    data = json.dumps({
        "name":"白内障",
        "disease_kind":"眼科",
        "info":
            '''
            老化，遗传、局部营养障碍、免疫与代谢异常，外伤、中毒、辐射等，都能引起晶状体代谢紊乱，导致晶状体蛋白质变性而发生混浊，称为白内障，此时光线被混浊晶状体阻扰无法投射在视网膜上，导致视物模糊。
            ''',
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