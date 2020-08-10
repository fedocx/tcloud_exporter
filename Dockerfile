FROM alpine:3.12.0
WORKDIR /data
COPY tencent_exporter /data/tencent_exporter
COPY config/metrics.yaml  /data/config/metrics.yaml
COPY config/tencent.yaml /data/config/tencent.yaml
CMD /data/tencent_exporter