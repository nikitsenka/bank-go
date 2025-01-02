FROM golang:alpine AS builder

RUN apk add --no-cache autoconf automake clang libpq-dev libtool make unixodbc-dev wget
WORKDIR /
RUN wget https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-17_00_0004-mimalloc.tar.gz
RUN tar -xvf REL-17_00_0004-mimalloc.tar.gz
WORKDIR /psqlodbc-REL-17_00_0004-mimalloc
RUN ln -sf /usr/lib/libpq.so.5 /usr/lib/libpq.so
RUN autoreconf -i && ./configure CC=clang CXX=clang++
RUN make && make install && make maintainer-clean

WORKDIR /app

COPY . .

ARG TARGET_ARCH=amd64

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.version=$VERSION

RUN echo Building for ${TARGET_ARCH}
RUN go env && go version
RUN GOOS=linux GOARCH=${TARGET_ARCH} \
go build -o bin/bank ./bank/

FROM alpine:latest

RUN apk add --no-cache \
  libpq \
  psqlodbc \
  unixodbc

COPY --from=builder /app/bin /app/bin
COPY tds.drive.template /app/tds.drive.template

RUN odbcinst -i -d -f /app/tds.drive.template

ENV PORT 8080
CMD ["/app/bin/bank"]
