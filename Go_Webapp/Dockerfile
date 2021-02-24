FROM golang:1.15.8-alpine3.13
RUN mkdir /task1
ADD . /task1
WORKDIR /task1
RUN go build -o main .
CMD ["/task1/main"]

