FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/xml2mysql ./cmd/xml2mysql
FROM alpine:3.21
RUN addgroup -S app && adduser -S -G app app
COPY --from=builder /out/xml2mysql /usr/local/bin/xml2mysql
USER app
ENTRYPOINT ["xml2mysql"]
