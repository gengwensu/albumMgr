
FROM golang:1.8-alpine

RUN mkdir -p /go/src/github.com/gengwensu/albumMgr

ADD . /go/src/github.com/gengwensu/albumMgr

WORKDIR /go/src/github.com/gengwensu/albumMgr

RUN go install github.com/gengwensu/albumMgr
# RUN apk add --no-cache curl

ENTRYPOINT /go/bin/albumMgr

EXPOSE 8081