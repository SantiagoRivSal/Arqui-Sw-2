FROM golang:1.17

ADD . /backend

WORKDIR /backend
RUN go mod init properties
RUN go mod tidy
RUN go build -o properties .
RUN chmod +x /backend

ENTRYPOINT ["/backend/properties"]