## http or https proxy

this is proxy for http or https with golang

### docker

```
docker run -d -p 8080:8080 pfinal/proxy
```

### example

```
curl --proxy http://127.0.0.1:8080 https://www.baidu.com
```


### build

```
go build
```


### run

```
./proxy --version

./proxy

./proxy --port :8080
```
