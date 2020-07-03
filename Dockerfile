FROM golang:1.14

ENV GO111MODULE=on
ENV WORKDIR /go/src/github.com/khanamoto/dokodemo
WORKDIR $WORKDIR

# go.mod go.sum が更新されたときのみレイヤを再構築する
COPY Makefile go.mod go.sum ./
RUN go mod download

RUN make setup

COPY . $WORKDIR

CMD ["realize", "start"]
