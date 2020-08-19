#!/bin/bash

chmod +x tcloud_exporter
docker build . -t tcloud_exporter:v1.0.0