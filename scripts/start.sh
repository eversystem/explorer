#!/bin/bash

# source ~/.bash_rc

# echo "Installing repository"
# git clone https://github.com/eversystem/explorer.git

echo " --------------------"
echo "|  Building frontend |"
echo " --------------------"

cd ~/explorer/frontend

echo "Building the server"
yarn
yarn build


echo " --------------------"
echo "|  Starting containers |"
echo " --------------------"

cd ~/explorer/

docker-compose up -d

