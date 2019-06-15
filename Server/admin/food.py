import requests
import json
from utils.config import *
ss = requests.session()
def addFood():
    data = json.dumps({
        "name":"马铃薯",
        # "food_kind_id":1,
        # "food_kind":"水果",
        "info":'''   
                  马铃薯，属茄科一年生草本植物，块茎可供食用，是全球第四大重要的粮食作物，仅次于小麦、稻谷和玉米。马铃薯又称地蛋、土豆 、洋山芋等，茄科植物的块茎。与小麦、稻谷、玉米、高粱并称为世界五大作物。
  马铃薯原产于南美洲安第斯山区，人工栽培历史最早可追溯到大约公元前8000年到5000年的秘鲁南部地区。
  马铃薯主要生产国有中国、俄罗斯、印度、乌克兰、美国等。中国是世界马铃薯总产最多的国家。
  2015年，中国将启动马铃薯主粮化战略，推进把马铃薯加工成馒头、面条、米粉等主食，马铃薯将成稻米、小麦 、玉米外的又一主粮。
                ''',
        "effect":"",
        "keyword":"",
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