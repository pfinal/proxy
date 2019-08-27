#FROM golang:1.12-alpine AS build-env
#WORKDIR /app
#ADD . /app
#ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
#RUN cd /app && go build -ldflags -s -a -installsuffix cgo main.go && ls -lh
FROM alpine:3.10
WORKDIR /app
#COPY --from=build-env /app/proxy /app
COPY ./proxy /app
EXPOSE 8080
CMD ["/app/proxy"]

# macOS
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o proxy ./main.go
# docker build  -t pfinal/proxy .
# docker run -d --restart=always --name proxy --cpus 0.1 -m 50m -p 8080:8080 pfinal/proxy
# curl --proxy http://127.0.0.1:8080 https://www.baidu.com
