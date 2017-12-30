#! /bin/bash
IM="golang:1.9.2-stretch"
CN="goenv"
CHN="os"

#docker container stop $CN
#docker container rm $CN

#LV="$(dirname $(pwd))"
LV="$(pwd)"        # local folder
CV="/root/repo/"   # container folder

P1="3000"           # expose port
P2="4000"           # expose port

docker run -it \
    -v $LV:$CV -p $P1:$P1 -p $P2:$P2 \
    --name $CN -h $CHN \
    $IM "/bin/bash"
