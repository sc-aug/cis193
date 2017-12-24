#! /bin/bash
IM="golang:1.9.2-stretch"
CN="goenv"
CHN="os"

#docker container stop $CN
#docker container rm $CN

#LV="$(dirname $(pwd))"
LV="$(pwd)"
CV="/root/repo/"

docker run -it \
    -v $LV:$CV \
    --name $CN -h $CHN \
    $IM "/bin/bash"
