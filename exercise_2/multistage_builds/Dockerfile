####
# Build container

FROM golang AS builder

WORKDIR /build

COPY main.go go.mod ./

RUN go build


####
# Runtime container

FROM alpine

ENV LISTEN_PORT=80

COPY --from=builder /build/app /usr/bin/app

CMD ["/usr/bin/app"]

