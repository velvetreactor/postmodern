FROM golang:1.10-alpine

RUN mkdir -p /go/src/github.com/velvetreactor/postapocalypse

WORKDIR /go/src/github.com/velvetreactor/postapocalypse

# Copy dep files
COPY ./docker.postapoc.src ./

# Install Go deps
RUN apk update && \
  apk add curl git && \
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
  go get github.com/codegangsta/gin && \
  dep ensure

CMD ["gin", "run", "main.go"]
