FROM golang:1.19-alpine

WORKDIR /go/src

RUN apk update && apk add git
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]
