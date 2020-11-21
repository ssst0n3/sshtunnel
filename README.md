# ssh tunnel

default port: 13400

## quick-to-start

### Method 1: docker
docker-compose up -d

docker-compose.yaml
```
version: '3'
services:
  sshtunnel:
    container_name: sshtunnel
    image: ssst0n3/sshtunnel:latest
    restart: always
    ports:
      - "13400:13400"
    environment:
      - SSH_IP=10.0.1.2
      - SSH_PORT=22
      - SSH_USERNAME=ctf
      - SSH_PASSWORD=SqjaVwE9
      - REMOTE_BIND_ADDRESS=127.0.0.1
      - REMOTE_BIND_PORT=13400
      - LOCAL_IP=0.0.0.0
      - LOCAL_PORT=13400
```

### Method 2: binary

download pre-build binary from [release](https://github.com/ssst0n3/sshtunnel/releases) 
or build from source code 
```
git clone https://github.com/ssst0n3/sshtunnel.git
cd sshtunnel/cli
GO111MODULE="on" GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o sshtunnel
```

run
```
SSH_IP=10.0.1.2 SSH_PORT=22 SSH_USERNAME=ctf SSH_PASSWORD=SqjaVwE9 REMOTE_BIND_ADDRESS=127.0.0.1 REMOTE_BIND_PORT=13400 LOCAL_IP=0.0.0.0 LOCAL_PORT=13400 ./sshtunnel
```

## repository
dockerhub: https://hub.docker.com/r/ssst0n3/sshtunnel
github: https://github.com/ssst0n3/sshtunnel

## golang usage
```
import "github.com/ssst0n3/sshtunnel"
```
 
