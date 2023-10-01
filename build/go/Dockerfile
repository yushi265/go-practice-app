# goバージョン
FROM golang:1.21.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
# パッケージのインポート
RUN go install golang.org/x/tools/cmd/goimports@latest
