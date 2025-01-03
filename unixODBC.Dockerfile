FROM golang:alpine AS builder

RUN apk add --no-cache autoconf automake clang libpq-dev libtool make unixodbc-dev wget

WORKDIR /app

COPY . .

ARG CC=gcc
ARG CXX=g++
ARG VERSION
ARG TARGET_ARCH=amd64

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.version=$VERSION

RUN echo Building for ${TARGET_ARCH}
RUN go env && go version
RUN GOOS=linux GOARCH=${TARGET_ARCH} \
CC=${CC} CXX=${CXX} \
go build -o bin/bank ./bank/

FROM alpine:latest

RUN apk add --no-cache \
  libpq \
  psqlodbc \
  unixodbc

COPY --from=builder /app/bin /app/bin
COPY <<EOT /tmp/tds.drive.template2
[PostgreSQL]
Description	= Official native client
Driver		  =/usr/lib/psqlodbca.so
Setup		    =/usr/lib/psqlodbcw.so
EOT
RUN odbcinst -i -d -f /tmp/tds.drive.template2

ENV PORT 8080
CMD ["/app/bin/bank"]
