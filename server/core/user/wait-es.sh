#!/bin/bash

# Kiểm tra sẵn sàng của Elasticsearch
echo "Waiting for Elasticsearch to be ready..."
while ! curl -s http://elasticsearch:9200 >/dev/null; do
    sleep 1
done
echo "Elasticsearch is ready"

# Kiểm tra sẵn sàng của MySQL
echo "Waiting for MySQL to be ready..."
while curl -s http://mysql:3306 >/dev/null; do
    sleep 1
done
echo "MySQL database is ready"

./core-user
