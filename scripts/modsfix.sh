#!/bin/bash

echo "Moving to root dir"

cd $(git rev-parse --show-toplevel)

echo "Starting hacky changes"

find . -type f -name "*.go" | xargs sed -i 's/\"code\.it4i\.cz\/lexis\/wp8\/dataset\-management\-interface\//\"code\.it4i\.cz\/lexis\/wp8\/dataset\-management\-interface\.git\//g'
find . -type f -name "*.go" | xargs sed -i 's/\"encoding\/json\"/\"github\.com\/segmentio\/encoding\/json\"/g'

echo "Hacky changes completed"

echo "Updating moduless..."

#go get code.it4i.cz/lexis/wp8/dataset-management-interface.git@connect_to_wp3_api

echo "Cleaning not necesary modules..."

go mod tidy

echo "Fix is completed"
