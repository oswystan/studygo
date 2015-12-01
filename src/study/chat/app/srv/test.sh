#!/bin/bash
while [ 1 ]
do
    echo ">>================"
    nc localhost 8000 <<EOF
wangyu
shenghuo
EOF
    echo "<<================"
    sleep 1
done
