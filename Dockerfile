# syntax=docker/dockerfile:1.7
FROM golang:1.26 AS builder

WORKDIR /go/src/github.com/Octops/agones-relay-http

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    make build && chmod +x /go/src/github.com/Octops/agones-relay-http/bin/agones-relay-http

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /go/src/github.com/Octops/agones-relay-http/bin/agones-relay-http /app/

ENTRYPOINT ["./agones-relay-http"]
