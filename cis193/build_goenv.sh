#! /bin/bash
IM="golang:1.9.2-stretch"
CN="goenv"
CHN="os"

#docker container stop $CN
#docker container rm $CN

#LV="$(dirname $(pwd))"
LV="$(pwd)"        # local folder
CV="/root/repo/"   # container folder
P="4000"           # expose port

docker run -it \
    -v $LV:$CV -p $P:$P \
    --name $CN -h $CHN \
    $IM "/bin/bash"
