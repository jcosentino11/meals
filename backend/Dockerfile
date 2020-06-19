# Based off of https://gist.github.com/afdolriski/c1252ba902f6c75ca14872a4e3d0074b#file-dockerfile
FROM golang:1.14-alpine

RUN apk add --update make

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

WORKDIR /dist

RUN cp /build/bin/main .

EXPOSE 8080

CMD ["/dist/main"]
