## 项目由来
考虑写这个项目主要有两个原因，一个是目前运行环境在腾讯云上，数据库使用得是pass数据库，但是腾讯云监控比较慢，对于pass数据库不便于安装探针。为了便于监控，我们采用Prometheus来实现监控数据的采集，所以就有编写exporter的想法。

### grafana 展示
![展示效果](image/grafana展示.png)
如要实现上图效果，导入[配置文件](https://github.com/fedocx/tcloud_exporter/blob/master/export_file/grafana/mysql%E7%9B%91%E6%8E%A7%E6%8C%87%E6%A0%87dashboard-1594188394862.json)
即可

## 支持列表

目前支持mysql数据库的监控数据采集，持续更新中，敬请期待。
