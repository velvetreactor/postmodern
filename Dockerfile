FROM golang:1.10-alpine

WORKDIR /go/src/app

# Copy dep files
COPY ./Gopkg* ./
COPY ./main.go ./
COPY ./src ./src

# Install Go deps
RUN apk update && \
  apk add curl git && \
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
  go get github.com/codegangsta/gin && \
  dep ensure

# Setup volume directories
RUN mkdir dist/

CMD ["gin", "run", "main.go"]
