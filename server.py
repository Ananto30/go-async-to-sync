from flask import Flask, request

import time
import requests
from multiprocessing import Pool

app = Flask(__name__)




def make_request(conversation_id):
    time.sleep(3)

    url = "http://localhost:8080/result"
    body = {
        "message": "hello world",
        "conversationId": conversation_id
    }
    requests.post(url, json=body)

    print("Callback URL called")


@app.route('/try-async', methods=['POST'])
def hello_world():

    body = request.json
    conversation_id = body['conversationId']

    pool = Pool(processes=1) # Start a worker processes.
    result = pool.apply_async(make_request, [conversation_id])

    return 'Hello, World!'



if __name__ == '__main__':
    app.run()