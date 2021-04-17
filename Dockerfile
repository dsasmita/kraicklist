FROM golang:1.16.3-alpine3.13
WORKDIR /root
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd