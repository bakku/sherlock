FROM golang:1.14-alpine as builder

WORKDIR /usr/app/src

COPY . .
RUN cd cmd/sherlock && go build

FROM alpine:3.9
COPY --from=builder /usr/app/src/cmd/sherlock/sherlock sherlock
CMD ["./sherlock"]
