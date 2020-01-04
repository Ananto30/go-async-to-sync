# Async to Sync
We often use legacy services that are designed to return result in callback urls and works in async styles. May be we only doing a query to fetch some results but that thing (response) is replied in a callback url we exposed. 

So I am trying here to make that thing synchronous, like a user can call `/async` endpoint to get the result directly rather opening a callback url. The Go server will handle those callback things. This code is a good example of how we can leverage Go's channel to do awesome things!

Here I am guessing the async server (Python) will expect a `trackId` with the request, so I created a `trackId` in the Go server and send it to the async server. It can work in the vice-versa way. May be this type of ID can be generated by the async server and pass it with the acknowledgement.

P.S. : If we want to scale this Go server, we cannot do so. Can you guess why?
Go channel cannot communicate between several servers right? So we can use  a simple pub-sub there to do the same thing, may be a Redis pub-sub, or RabbitMQ. I will make a new repo about that soon.

## Start async server
This server is written in python. It's in the `async-server-python` folder. 
You can create a venv if you wish with python 3.7(+)
```
cd async-server-python
python3 -m venv venv
source venv/bin/activate
```
Install the requirements -
`pip install -r requirements.txt`

Start the server -
`python server.py`

## Start the Go server
This server will do all the works to make that async request to sync. 
Install the dependencies (I used Dep as dependency manager) - 
`dep ensure`
Start the server - 
`go run main.go`

## Test it
Endpoint - [localhost:8080/async](localhost:8080/async) 
