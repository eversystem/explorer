#!/bin/bash

cd /go/src/github.com/iost-official/backend

rm explorer
make

cd ./task

rm explorer-task
make

cd /go/src/github.com/iost-official/backend

nohup ./explorer  &

nohup ./task/explorer-task  &

/bin/bash

