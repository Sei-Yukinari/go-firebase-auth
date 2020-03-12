#Step 1 Builld
FROM golang:alpine as builder

RUN apk update \
  && apk add git

WORKDIR /go/src/api/

COPY src/api/ .

RUN go mod download

RUN go build -o /server

#Step2 Move build file
FROM alpine:3.9

RUN apk add --no-cache     mysql-client \
                           mysql-dev \
                           alpine-sdk

COPY --from=builder /server .

ENTRYPOINT ["sh","/go/src/api/entrypoint.sh"]
