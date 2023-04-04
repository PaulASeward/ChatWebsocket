FROM golang:1.12-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN apk add --no-cache git
RUN apk add --no-cache build-base
RUN go mod download
RUN go build -o main ./main.go
CMD ["/app/main"]