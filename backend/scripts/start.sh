#!/bin/bash

cd /go/src/github.com/iost-official/backend

rm explorer
make

cd ./task

rm explorer-task
make

cd /go/src/github.com/iost-official/backend

export HOST="3.115.13.190"

nohup ./explorer  &

nohup ./task/explorer-task  &

/bin/bash

