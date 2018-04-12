FROM golang:alpine


ADD . /go/src/github.com/nearmap/cvm-example
RUN go install github.com/nearmap/cvm-example

RUN rm -r /go/src/github.com/nearmap/cvm-example

VOLUME /go/src

EXPOSE 8081

ENTRYPOINT ["cvm-example"]
