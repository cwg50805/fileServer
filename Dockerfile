# build stage
FROM golang:1.16-alpine
WORKDIR /go/src/hp/
COPY . .
EXPOSE 4300
CMD go run main.go

