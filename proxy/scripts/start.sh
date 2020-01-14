#!/bin/ash

export PATH=/usr/local/bin:$PATH 

cd /app/proxy

echo "installing forever"
npm i -g forever ts-node


echo "installing dependencies"
yarn

echo "Starting the proxy server"

forever start -c ts-node ./forever/config.json

ash