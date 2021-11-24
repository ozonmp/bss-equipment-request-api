# Builder

ARG GITHUB_PATH=github.com/ozonmp/bss-equipment-request-api

FROM golang:1.16-alpine AS builder
RUN apk add --update make git curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make build-migration-go

# migration

FROM alpine:latest as migration
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/migration .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root migration

CMD ["./migration"]
