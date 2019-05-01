FROM golang:1.8

WORKDIR /go/src/app

RUN go get -d -v github.com/nikitsenka/bank-go
RUN go install -v github.com/nikitsenka/bank-go

CMD ["bank-go"]