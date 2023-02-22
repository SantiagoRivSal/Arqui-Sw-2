FROM golang:1.17

ADD . /search

WORKDIR /search
RUN go mod init consumer
RUN go mod tidy
RUN go build -o consumer .
RUN chmod +x /search

ENTRYPOINT ["/search/consumer"]