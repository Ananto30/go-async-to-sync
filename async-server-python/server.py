from flask import Flask, request

import time
import requests
from multiprocessing import Pool

app = Flask(__name__)


# This is actually the client's webhook, where the async response will be sent
callback_url = "http://localhost:8005/result"


def make_request(track_id):
    """
    This is a mock async processor, which post the response to the registered webhook
    """
    # Mocking work
    time.sleep(3)

    # A demo response to the webhook
    response = {"message": "hello world", "trackId": track_id}
    # POST request to the webhook
    requests.post(callback_url, json=response)

    print(f"Callback URL called with track_id {track_id}")


# Start a worker processes
pool = Pool(processes=1)


@app.route("/try-async", methods=["POST"])
def try_async():
    """
    Request handler for async request, takes trackId in body
    """

    body = request.json
    # Parse the trackId from body
    track_id = body["trackId"]

    # This will make the async request to the callback url
    # and this is non-blocking
    pool.apply_async(make_request, [track_id])

    # Just a demo response, one may get acknowledgement of the async request
    return "Hello, World!"


if __name__ == "__main__":
    app.run()
