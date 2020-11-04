FROM golang:1.15.3-alpine
WORKDIR /ping-pong
ADD . /ping-pong
RUN cd /ping-pong && go build -o ping-pong
EXPOSE 9936
ENTRYPOINT ./ping-pong