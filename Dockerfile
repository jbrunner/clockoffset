FROM golang:1.25-alpine AS builder
WORKDIR /build
COPY go.mod go.sum *.go ./
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o clockoffset .

FROM scratch
COPY --from=builder /build/clockoffset /clockoffset
USER 65534:65534
ENTRYPOINT ["/clockoffset"]
