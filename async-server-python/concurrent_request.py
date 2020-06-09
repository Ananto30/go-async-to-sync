import requests
import time
from concurrent.futures import ThreadPoolExecutor


def async_req(url):
    body = {"accountId": "3"}
    res = requests.post(url, json=body)
    print(res.json())

urls = [
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",
    "http://localhost:8005/async",

]
with ThreadPoolExecutor(max_workers=10) as pool:
    list(pool.map(async_req,urls))