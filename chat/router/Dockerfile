FROM golang:latest
MAINTAINER carsonsx <carsonsx@qq.com>
COPY . /go/src/github.com/carsonsx/net4g-chat/router
RUN go install -ldflags "-s -w" github.com/carsonsx/net4g-chat/router
WORKDIR /go/bin
ENTRYPOINT ["/go/bin/router"]
EXPOSE 9000