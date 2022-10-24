#!/bin/bash

curl 127.0.0.1:9200/metadata -XDELETE

curl 127.0.0.1:9200/metadata -XPUT -d'{"mappings":{"objects":{"properties":{"name":{"type":"string","index":"not_analyzed"},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"string"}}}}}'

