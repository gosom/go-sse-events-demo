# Server sent events demo using golang

Example usage of how to monitor a redis pub sub channel and stream the results to clients.

Uses server sent events to stream the messages to the client.

Uses only one goroutine per client and only one connection to redis.

Also, it uses the http/2 protocol and it includes a self signed certificate.


**Notice**: This is only demo code to help a friend. It's barely tested


### Quickstart

The command below starts:
- a redis server
- a client that publishes messages to a redis channel
- a web server that listens on port 8080

```
docker-compose up
```

Then:

```
curl -k https://localhost:8080
```


you should see the incoming messages


