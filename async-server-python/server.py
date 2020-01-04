from flask import Flask, request

import time
import requests
from multiprocessing import Pool

app = Flask(__name__)


def make_request(track_id):
    time.sleep(3)

    callback_url = "http://localhost:8080/result"
    body = {
        "message": "hello world",
        "trackId": track_id
    }
    requests.post(callback_url, json=body)

    print("Callback URL called")


@app.route('/try-async', methods=['POST'])
def hello_world():

    body = request.json
    track_id = body['trackId']

    pool = Pool(processes=1) # Start a worker processes.
    # This will make the async request to the callback url
    result = pool.apply_async(make_request, [track_id])

    # Just a demo response, one may get acknowledgement of the async request
    return 'Hello, World!'



if __name__ == '__main__':
    app.run()