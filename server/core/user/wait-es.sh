#!/bin/bash

# Kiểm tra sẵn sàng của Elasticsearch
echo "Waiting for Elasticsearch to be ready..."
while ! curl -s http://elasticsearch:9200 >/dev/null; do
    sleep 1
done

echo "Elasticsearch is ready"

./core-user
