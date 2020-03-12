FROM golang:1.14-alpine as builder

WORKDIR /usr/app/src

COPY . .
RUN cd cmd/sherlock && go build

FROM alpine:3.9
COPY --from=builder /usr/app/src/cmd/sherlock/sherlock sherlock
COPY --from=builder /usr/app/src/web/templates templates

ENV TEMPLATE_DIR templates/
CMD ["./sherlock"]
