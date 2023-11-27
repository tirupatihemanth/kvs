#!/bin/bash
nginx -s stop
cur_dir=$(pwd)
conf_path="$cur_dir/conf.d/nginx.conf"
nginx -c "$conf_path"

for((n1=65431; n1<=$1; n1+=1)); do
    ./kvs $n1 &
    echo $!
done

echo "Don't forget to kill above PIDs. kill <pid>"
# sleep 1
# python client.py $2 $3

# con_ids=$(docker ps -q)
# for con_id in $con_ids; do
#     docker stop "$con_id"
# done