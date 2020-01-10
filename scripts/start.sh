#!/bin/bash

source ~/.bash_rc

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

echo "Building the server"
make

echo "Starting the server"

nohup ./explorer &


echo "Starting cron task"

cd ~/explorer/backend/task

echo "Building the server"

make

echo "Starting the server"

nohup ./explorer-task &


echo " --------------------"
echo "|  Starting frontend |"
echo " --------------------"

cd ~/explorer/frontend

echo "Installing dependencies"
npm install -g forever
yarn

echo "Building the server"
yarn build


echo "Starting the server"
docker run -d -p 8080:80 -v ~/explorer/frontend/dist/://usr/share/nginx/html:ro nginx
