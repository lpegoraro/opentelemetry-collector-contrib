FROM golang:1.21-alpine as builder

RUN apk --update add ca-certificates git tzdata

WORKDIR /build

RUN go install go.opentelemetry.io/collector/cmd/builder@latest
COPY opentelemetry-collector-builder.yaml .
RUN builder --config=opentelemetry-collector-builder.yaml


FROM scratch

ARG USER_UID=10001

USER ${USER_UID}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# this may need to be customized based on opentelemetry-collector-builder.yaml config
COPY --from=builder /tmp/otelcol/otelcol /otelcol
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY example-config.yaml /config.yaml

ENTRYPOINT ["/otelcol"]
CMD ["--config", "/config.yaml"]