FROM golang

ENV LISTEN_PORT=80

WORKDIR /app

COPY main.go go.mod ./

RUN go build

RUN mv /app/app /usr/bin/app

CMD ["/usr/bin/app"]
