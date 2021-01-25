FROM golang:alpine
RUN apk add git
WORKDIR /go/src/github.com/jbrunner/clockoffset/
RUN go get github.com/beevik/ntp
COPY clockoffset.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o clockoffset .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=0 /go/src/github.com/jbrunner/clockoffset/clockoffset .
ENTRYPOINT ["./clockoffset"] 
