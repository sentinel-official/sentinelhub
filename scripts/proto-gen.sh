#!/usr/bin/env bash

set -euo pipefail

project_dir=$(pwd)

cd "${project_dir}/third_party/osmosis/proto/" && \
buf generate && \
cd ../ && \
cp -r github.com/sentinel-official/sentinelhub/v12/third_party/osmosis/* ./ && \
rm -rf github.com/

cd "${project_dir}/proto/" && \
buf generate && \
cd ../ && \
cp -r github.com/sentinel-official/sentinelhub/v12/* ./ && \
rm -rf github.com/
