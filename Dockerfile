FROM golang:1.12

ARG LOG_DIR=/app/logs
VOLUME ["/app/logs"]

RUN mkdir -p ${LOG_DIR}

VOLUME ["/go/src/app"]
WORKDIR $GOPATH/src/app

COPY run.sh /bin/run.sh
