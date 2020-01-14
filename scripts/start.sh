#!/bin/bash

source ~/.bash_rc

echo "Installing repository"
git clone https://github.com/eversystem/explorer.git


echo " --------------------"
echo "|  Starting backend  |"
echo " --------------------"

echo "Installing dependencies"

cd ~/explorer/backend
go mod init github.com/iost-official/explorer/backend

echo "Building the explorer server"
make

cd ~/explorer/backend/task

echo "Building cron task"

make


echo " --------------------"
echo "|  Building frontend |"
echo " --------------------"

echo "Building the server"
yarn build


echo " --------------------"
echo "|  Starting containers |"
echo " --------------------"

cd ~/explorer/

docker-compose up -d

