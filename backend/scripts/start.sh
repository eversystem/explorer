#!/bin/bash

cd /go/src/github.com/iost-official/backend

rm explorer
make

cd ./task

rm explorer-task
make

cd /go/src/github.com/iost-official/backend

nohup ./explorer  &

cd task 

nohup ./explorer-task  &

/bin/bash

