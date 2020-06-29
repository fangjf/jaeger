#!/usr/bin/env bash

exec env SPAN_STORAGE_TYPE=elasticsearch cmd/query/query-darwin-amd64 --es.server-urls=http://172.21.3.152:9200