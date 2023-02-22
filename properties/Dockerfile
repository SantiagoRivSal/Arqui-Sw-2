FROM golang:1.17

ADD . /properties

WORKDIR /properties
RUN go mod init properties
RUN go mod tidy
RUN go build -o properties .
RUN chmod +x /properties

ENTRYPOINT ["/properties"]