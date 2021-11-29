# Builder

ARG GITHUB_PATH=github.com/ozonmp/bss-equipment-request-api

FROM golang:1.16-alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make deps-go
RUN make build-retranslator-go

# Retranslator

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/retranslator .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .

RUN chown root:root retranslator

EXPOSE 9103

CMD ["./retranslator"]
