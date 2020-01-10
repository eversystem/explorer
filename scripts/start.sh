#!/bin/bash

echo "Installing repository"
git clone https://github.com/eversystem/explorer.git


echo " --------------------"
echo "|  Starting mongodb  |"
echo " --------------------"

sudo service mongod start

echo " --------------------"
echo "|  Starting backend  |"
echo " --------------------"


echo "Installing dependencies"

cd ~/explorer/backend
go mod init github.com/iost-official/explorer/backend

"Building the server"
make

"Starting the server"

nohup ./explorer &


echo "Starting cron task"

cd ~/explorer/backend/task

"Building the server"

make

"Starting the server"

nohup ./explorer-task &


echo " --------------------"
echo "|  Starting frontend |"
echo " --------------------"

cd ~/explorer/frontend

echo "Installing dependencies"
npm install -g forever
yarn

"Building the server"
yarn build


"Starting the server"
docker run docker run -d -p 8080:80 -v ~/explorer/frontend/dist/://usr/share/nginx/html:ro nginx
