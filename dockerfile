FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN cd ./cmd/main
RUN go build -o main ../../bin
CMD ["/app/bin/main"]