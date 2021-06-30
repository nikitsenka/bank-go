FROM golang as builder

WORKDIR /app

COPY . .

ARG TARGET_ARCH=amd64

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.version=$VERSION

RUN echo Building for ${TARGET_ARCH}
RUN go env && go version
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGET_ARCH} \
	go build -o bin/bank ./bank/

FROM scratch
COPY --from=builder /app/bin /app/bin
ENV PORT 8080
ENTRYPOINT ["/app/bin/bank"]
