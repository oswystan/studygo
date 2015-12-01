#!/bin/bash
while [ 1 ]
do
    nc localhost 8000 >/dev/null 2>&1<<EOF
wangyu
shenghuo
bye
EOF
    if [ $? -ne 0 ]; then
        echo `date +'%Y-%m-%d %T'` "server down!!"
    else
        echo `date +'%Y-%m-%d %T'` "server up"
    fi
    sleep 1
done
