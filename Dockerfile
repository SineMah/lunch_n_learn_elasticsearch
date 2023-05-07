FROM golang:alpine

RUN apk add --no-cache \
  curl \
  zip

RUN mkdir /app

WORKDIR /app

CMD ["tail", "-f", "/dev/null"]