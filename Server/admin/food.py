import requests
import json
from utils.config import *
ss = requests.session()
def addFood():
    data = json.dumps({
        "name":"木耳",
        # "food_kind_id":1,
        # "food_kind":"水果",
        "info":'''   
                木耳主要生长在中国和日本。中国大部分是东北木耳和秦岭木耳。木耳既可野生又可以人工培植，种子实体呈耳状、叶状或杯状、薄、边缘波浪状，宽3-10厘米，厚2毫米左右，以侧生的短柄或狭细的附着部固着于基质上。色泽黑褐，质地柔软呈胶质状，薄而有弹性，湿润时半透明，干燥时收缩变为脆硬的角质近似革质。味道鲜美，可素可荤，营养丰富。木耳味甘，性平，具有很多药用功效。能益气强身，有活血效能，并可防治缺铁性贫血等；可养血驻颜，令人肌肤红润，容光焕发，能够疏通肠胃，润滑肠道，同时对高血压患者也有一定帮助 ''',
        "effect":"",
        "keyword":"补气养血;润肺止咳;止血;降压;抗癌",
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