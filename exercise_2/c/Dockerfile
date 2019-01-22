FROM alpine

ENV VARIABLE=1

RUN apk add --no-cache \
      gcc \
      musl-dev

WORKDIR /build

COPY app.c .

RUN gcc app.c -o app

RUN rm /build/app.c

ENTRYPOINT ["/build/app"]
