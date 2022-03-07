FROM golang:alpine
WORKDIR /go/src/app
COPY main.go .
RUN go mod init
RUN go mod tidy
RUN go build 

CMD ["./app"]