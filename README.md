## http or https proxy

this is proxy for http or https with golang

### build

```
go build
```


### run

```
./proxy

./proxy --port :8080

docker run -d --restart=always --name proxy --cpus 0.1 -m 50m -p 8080:8080 pfinal/proxy

```

### example

```
curl --proxy http://127.0.0.1:8080 https://www.baidu.com
```
