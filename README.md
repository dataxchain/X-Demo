# dataplatform


### Hyperledger Fabric 네트워크 설치

- Docker Compose로 되어있음
- TLS 관련 기능을 끄기 위하여 설정을 변경함

```
$ cd ./app/fabric/network
$ ./byfn.sh up
```

### assetbox 체인코드 설치 및 테스트

- 예제에서 기본 제공되는 체인코드 설치, 실행 스크립트를 assetbox 체인코드에 맞게 수정하였음

```
$ sudo docker exec cli scripts/assetbox_script.sh
```

### SDK 실행을 위한 node.js 컨테이너 실행

```
$ sudo docker-compose up
$ sudo docker exec -it nodeapp bash

root@c0e4a10cf951:/usr/src/app# node middleware/enrollAdmin.js
root@c0e4a10cf951:/usr/src/app# node middleware/registerUser.js
root@c0e4a10cf951:/usr/src/app# node web/app.js
```