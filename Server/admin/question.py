import requests
import json
from utils.config import *
ss = requests.session()
def addQuestion():
    data = json.dumps({
        "info":"你容易健忘吗",
        "answer_a":"是",
        "response_a":"健忘" ,
        "answer_b":"否",
        "response_b":"不健忘" ,
        "answer_c":"不知道" ,
        "response_c":"",
    })
    url = URL.format(env=cur_url, term=admin_url, action="/addQuestion")
    postReq(url,ss,data,admin_token)




def main():
    addQuestion()
if __name__ == '__main__':
    main()