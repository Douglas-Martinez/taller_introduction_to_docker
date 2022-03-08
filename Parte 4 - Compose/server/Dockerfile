FROM golang:alpine
WORKDIR /go/src/app
COPY main.go .
RUN go mod init
RUN go mod tidy 
CMD ["go","run","main.go"]

# Producción
# RUN go build
# CMD ["./app"]