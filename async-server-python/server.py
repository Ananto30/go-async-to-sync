from flask import Flask, request

import time
import requests
from multiprocessing import Pool

app = Flask(__name__)


# This is actually the client's webhook, where the async response will be sent
callback_url = "http://localhost:8005/result"

"""
data format - 
{
    accountId: balance
}
"""
demo_balance_data = {"1": "1000", "2": "4377", "3": "9999999", "12345": "12345"}


def make_request(track_id, account_id):
    """
    This is a mock async processor, which post the response to the registered webhook
    """
    # Mocking extensive work
    time.sleep(3)

    balance = demo_balance_data.get(account_id, None)
    response = {"trackId": track_id, "balance": balance}
    # POST request to the webhook
    requests.post(callback_url, json=response)

    print(f"Callback URL called with track_id {track_id}")


# Start some worker processes
pool = Pool(processes=4)


@app.route("/async-balance", methods=["POST"])
def try_async():
    """
    Request handler for async request, takes trackId in body
    """

    body = request.json
    # Parse the trackId and accountId from body
    print(body)
    track_id = body["trackId"]
    account_id = body["accountId"]

    # This will make the async request to the callback url
    # and this is non-blocking
    pool.apply_async(make_request, [track_id, account_id])

    # Just a demo response, one may get acknowledgement of the async request
    return "Hello, World!"


if __name__ == "__main__":
    app.run()
