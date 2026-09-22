FROM golang:1.24-alpine AS builder
WORKDIR /src

COPY go.mod ./
COPY . .

# v0.1 does not yet commit go.sum. Resolve and verify module metadata
# inside the isolated builder stage before compiling.
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/xml2mysql ./cmd/xml2mysql

FROM alpine:3.21
RUN addgroup -S app && adduser -S -G app app
COPY --from=builder /out/xml2mysql /usr/local/bin/xml2mysql
USER app
ENTRYPOINT ["xml2mysql"]
