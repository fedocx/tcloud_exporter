FROM alpine:3.12.0
WORKDIR /data
COPY tcloud_exporter /data/tcloud_exporter
COPY config/metrics.yaml  /data/config/metrics.yaml
COPY config/tencent.yaml /data/config/tencent.yaml
CMD /data/tcloud_exporter