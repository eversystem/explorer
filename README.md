# Explorer

## Prerequisites

OS:　Amazon Linux2
Docker, Docker Composeがインストール済みでDockerデーモンが起動していること

エクスプローラーを建てるサーバーは以下をセキュリティグループで開ける

1. Port: 8080, source: 0.0.0.0/0
2. Port: 8000, source: 0.0.0.0/0

### 手順１

以下のスクリプトで必要なものをインストールする

```

./scripts/init_al2.sh

```

### 手順2

各種コンフィグを手で書き換える

1. ./docker-compose.ymlのenvironment
   1. 以下を書き換える（docker-compose.yml内）
   2. backend/HOST
   3. proxy/BASICNAME, BASICPASS, JWTSEC
2. backendのコンフィグ
   1. ./backend/config/config.json
      1. rpcHOSTをノードのIPに
3. frontendのコンフィグ
   1. ./frontend/config/prod.env.js
      1. APIの部分にノードのIPを追加

### 手順3 

以下のスクリプトでコンテナを起動する

```

./scripts/start.sh

```

### 手順4

バックエンドを起動する
<!-- TODO 自動化 -->

```

docker exec -it back bash

/go/src/github.com/iost-official/backend
に入るので

nohup ./explorer &

cd ./task

nohup ./explorer-task & 

```

### 手順5

サーバーのIP:8080
に行って入れるか確認

