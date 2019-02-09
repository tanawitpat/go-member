ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine as base

EXPOSE 8050
WORKDIR /go/src/go-member
COPY . .
RUN go build -o main .

FROM base as dummy

CMD ["/go/src/go-member/main"]