#!/bin/bash

cd /go/src/github.com/iost-official/backend

rm explorer
make

cd ./task

rm explorer-task
make

cd /go/src/github.com/iost-official/backend

nohup ./explorer 2>&1 &

nohup ./task/explorer-task 2>&1 &

/bin/bash

