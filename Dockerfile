FROM golang:1.11-stretch

WORKDIR /go/src/github.com/Candy_Wheel_Alert

ENV ENV LOCAL
ARG secret
ENV SECRET ${secret}
ARG token
ENV TOKEN ${token}
ARG id
ENV ID ${id}
ENV PACE */60 * * * * *

ADD . .

CMD ["/usr/local/go/bin/go", "run", "main.go"]