FROM golang:1.18

ADD . /users

WORKDIR /users
RUN go mod init users
RUN go mod tidy
RUN go build -o users .
RUN chmod +x /users

<<<<<<< HEAD:users/dockerfile
ENTRYPOINT ["/users/users"]
=======
ENTRYPOINT ["/users"]
>>>>>>> 2a7239ac1f0d2329ea95226d2a873900198c641a:users/Dockerfile
