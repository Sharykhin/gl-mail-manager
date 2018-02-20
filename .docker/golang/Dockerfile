FROM golang:1.9

ENV TEST_FAIL NO

ADD . /go/src/github.com/Sharykhin/gl-mail-manager

WORKDIR /go/src/github.com/Sharykhin/gl-mail-manager

RUN go get .

RUN go install github.com/Sharykhin/gl-mail-manager

ENTRYPOINT /go/bin/gl-mail-manager

#CMD tail -f /dev/null