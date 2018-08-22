FROM instrumentisto/dep as dep
WORKDIR /go/src/github.com/ruelephant/MailRuTestApp
ADD . /go/src/github.com/ruelephant/MailRuTestApp
RUN dep ensure -vendor-only

FROM golang as builder
WORKDIR /go/src/github.com/ruelephant/MailRuTestApp
COPY --from=dep /go/src/github.com/ruelephant/MailRuTestApp ./
WORKDIR /go/src/github.com/ruelephant/MailRuTestApp/cmd/main
RUN go build -o /webapp/webserver .
WORKDIR /webapp
CMD ./webserver