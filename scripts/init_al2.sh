#!/bin/bash

## A script to install explorer dependencies into amazon linux2 instance
## pleade install docker manually beforehand

echo "Updating package"
sudo yum update

echo "Installing dependencies"
sudo yum -y install gcc gcc-c++ make
sudo yum install git

echo "Installing go"
git clone https://github.com/syndbg/goenv.git ~/.goenv

export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH 
eval "$(goenv init -)"

goenv install 1.11.4
goenv global 1.11.4
goenv rehash

echo "go version"
go version


echo "Installing node JS"

curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.2/install.sh | bash

echo 'export NVM_DIR="$HOME/.nvm"' >> ~/.bash_rc
echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"' >> ~/.bash_rc

source ~/.bash_rc

nvm install lts/erbium
nvm use lts/erbium

echo "node version"
node -v

npm install -g yarn


# Installing MongoDB

sudo touch /etc/yum.repos.d/mongodb-org-4.2.repo

sudo -i

echo "[mongodb-org-4.2]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/amazon/2/mongodb-org/4.2/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.2.asc" > /etc/yum.repos.d/mongodb-org-4.2.repo

exit

sudo yum install -y mongodb-org


echo "Notice: please export an environment variable HOST"
echo "HOST should be v4 IP address of an IOST node"

