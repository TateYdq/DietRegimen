import requests
import json
from utils.config import *
ss = requests.session()

def getQuestionnaire():
    data = {
    }
    url = URL.format(env=cur_url, term=client_url, action="/recommend/getQuestionnaire")
    getReq(url, ss, data=data,token=user_token)

def submitQuestionnaire():
    data =json.dumps({
        "answer_sheets":[
            {
                "question_id":1,
                "answer":"不头晕",
            },
            {
                "question_id": 2,
                "answer": "不眼花",
            }
        ]
    })
    url = URL.format(env=cur_url, term=client_url, action="/recommend/submitQuestionnaire")
    postReq(url, ss, data, token=user_token)

def main():
    # getQuestionnaire()
    submitQuestionnaire()

if __name__ == '__main__':
    main()