# go開発環境
FROM golang

WORKDIR /go/src/app
COPY . .

ENV http_proxy "http://sg-ss16-1.isid.co.jp:8080"
ENV https_proxy "http://sg-ss16-1.isid.co.jp:8080"

RUN apt-get install git 
RUN go get github.com/pilu/fresh
CMD ["fresh"]