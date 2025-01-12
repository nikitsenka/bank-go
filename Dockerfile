FROM golang AS builder

RUN apt-get update \
&& export DEBIAN_FRONTEND=noninteractive \
&& apt install -y automake build-essential clang iodbc libiodbc2-dev libodbc2 libpq-dev libtool odbc-postgresql \
&& apt clean -y \
&& rm -rf /var/lib/apt/lists/*

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
RUN ln -s /usr/lib/x86_64-linux-gnu/libodbc.so.2.0.0 /usr/lib/x86_64-linux-gnu/libodbc.so
RUN CGO_ENABLED=1 GOOS=linux GOARCH=${TARGET_ARCH} \
CC=${CC} CXX=${CXX} \
CGO_CFLAGS="-I/usr/include/iodbc" \
go build -o bin/bank ./bank/

FROM debian:stable-slim

RUN apt-get update \
&& export DEBIAN_FRONTEND=noninteractive \
&& apt install -y iodbc libodbc2 libpq5 odbc-postgresql \
&& apt clean -y \
&& rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin /app/bin
COPY <<EOT /tmp/tds.drive.template2
[PostgreSQL]
Description	= Official native client
Driver		  =/usr/lib/x86_64-linux-gnu/odbc/psqlodbca.so
Setup		    =/usr/lib/x86_64-linux-gnu/odbc/psqlodbcw.so
EOT
RUN odbcinst -i -d -f /tmp/tds.drive.template2

ENV PORT 8080
CMD ["/app/bin/bank"]
