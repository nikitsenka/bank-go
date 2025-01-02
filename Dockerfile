FROM golang:alpine AS builder

RUN apk add --no-cache autoconf automake clang libpq-dev libtool make unixodbc-dev wget
WORKDIR /
RUN wget https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-17_00_0004-mimalloc.tar.gz
RUN tar -xvf REL-17_00_0004-mimalloc.tar.gz
WORKDIR /psqlodbc-REL-17_00_0004-mimalloc
RUN ln -sf /usr/lib/libpq.so.5 /usr/lib/libpq.so
RUN autoreconf -i && ./configure CC=clang CXX=clang++
RUN make && make install && make maintainer-clean && cd / && rm -rf /psqlodbc-REL-17_00_0004-mimalloc

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
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/tds.drive.template /app/tds.drive.template
COPY --from=builder /usr/local/lib /usr/local/lib

RUN apk add --no-cache libpq unixodbc
RUN odbcinst -i -d -f /app/tds.drive.template
COPY <<"EOT" /etc/odbc.ini
[ODBC]
Driver = PostgreSQL
Description = PostgreSQL Data Source
Servername = localhost
Port = 5432
Protocol = 11.2
UserName = postgres
Password = test1234
Database = postgres
EOT
ENV PORT 8080
CMD ["/app/bin/bank"]
