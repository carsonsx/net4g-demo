FROM golang:latest
MAINTAINER carsonsx <carsonsx@qq.com>
COPY . /go/src/github.com/carsonsx/net4g-chat/server
RUN go install -ldflags "-s -w" github.com/carsonsx/net4g-chat/server
WORKDIR /go/bin
ENTRYPOINT ["/go/bin/server"]
EXPOSE 8000