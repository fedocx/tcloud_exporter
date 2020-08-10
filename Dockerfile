FROM alpine:3.12.0
WORKDIR /data
COPY tencent_exporter /data/tencent_exporter
COPY metrics.yaml  /data/config/metrics.yaml
COPY tencent.yaml /data/config/tencent.yaml
CMD tencent_exporter