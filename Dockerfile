FROM golang:1.6-onbuild

WORKDIR /go/src/github.com/otobrglez/rss-machine

ADD . /go/src/github.com/otobrglez/rss-machine

RUN go get github.com/tools/godep && godep restore && godep go build && godep go install

EXPOSE 5000

CMD ["rss-machine"]
