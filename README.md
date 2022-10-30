# webhook-test-server
Simple server for testing webhooks

```sh
docker build -t webhook-test-server .
```

```sh
docker run -it -p 80:8080 webhook-test-server -u=/hooks/123
```

An easy way of coming up with an "anonymous" endpoint - is to use the following 

```sh
md5 -s "Some random text"
```
Which will generate a string - in this case `f15c7ac831170f4a1fc5512c1a4c7acf`.