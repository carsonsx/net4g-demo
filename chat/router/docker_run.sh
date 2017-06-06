#!/usr/bin/env bash

prog=chatrouter
originport=9000
exportport=9000

docker ps -a | grep $prog | awk '{print $1}' | xargs docker rm -f

num=$1

if [ -z $num ]; then
    docker run -itd --net mynet --name $prog -p $exportport:$originport $prog
else
    for((i=1;i<=$num;i++))
    do
       let port=$exportport+$i
       docker run -itd --net mynet --name $prog$i -p $port:$originport $prog
    done
fi


