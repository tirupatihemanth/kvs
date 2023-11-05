#!/bin/bash
nginx -s stop
nginx

for((n1=65431; n1<=65431; n1+=1)); do
    echo "$(docker run -e PORT="$n1" -dp "$n1":"$n1" docker-kvs:v2.0)"
done

sleep 1
python client.py

con_ids = $(docker ps -q)
for con_id in $con_ids
    docker stop "$con_id"
done