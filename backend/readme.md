Iost explorer backend
======

j-LOD STG
  "rpcHost": "172.31.43.158:30002",



### Requirements 

1. golang 1.10+
2. mongodb



### How to deploy

1. clone code to $GOPATH/src/github.com/iost-official/explorer
2. add config file cp backend/config/config.json.sample backend/config/config.json
    update the ```rpcHost```, ```mongodb``` config if need
3. Run blockchain sync task: 
```bash
cd backend/task
make
nohup ./explorer-task &
```
4. Run REST api service
```bash
cd backend
make
nohup ./explorer &
```


