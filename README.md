## http or https proxy

this is proxy for http or https with golang

### run

```
curl -o proxy https://github.com/pfinal/proxy/releases/download/v1.0.0/proxy-linux

chmod +x proxy

./proxy --version

./proxy

./proxy --port :8080
```

### docker

```
docker run -d -p 8080:8080 pfinal/proxy
```

### example

```
curl --proxy http://127.0.0.1:8080 https://www.baidu.com
```

