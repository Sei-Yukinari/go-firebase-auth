#Step 1 ビルド処理のみ
FROM golang:alpine as builder

RUN apk update \
  && apk add git

WORKDIR /go/src/api/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /server

#Step2 ビルドしたファイルを実際に使うために移動
FROM alpine:3.9

COPY --from=builder /server .

ENTRYPOINT ["/server"]
