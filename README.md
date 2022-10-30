# webhook-test-server
Simple server for testing webhooks

```sh
docker build -t webhook-test-server .
```

```sh
docker run -it -p 80:8080 webhook-test-server --path /hooks/123
```