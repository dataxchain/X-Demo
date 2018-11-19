### 1. 개발

```
$ cd k8s
$ kubectl create -f goapp.yaml

$ telepresence --swap-deployment goapp --docker-run --rm -it \
-v ../../app/chaincode/src/dataplatform:/go/src/app \
DataXchain/chaincode:latest bash

# telepresence를 통해 다음과 같이 배포 환경과 동일한 환경에서 개발 중인 모듈 실행 가능
root@1baee5a0784f:/go/src/app# go test
```

- `src/dataplatform` 디렉토리가 컨테이너의 `/go/src/app`과 연결되어있으므로 코드 수정시 컨테이너에 바로 반영됨
- 리눅스 서버의 `src/dataplatform` 을 sftp로 IDE와 연동하면 편리하게 사용할 수 있음

### 2. 배포

- 개발이 끝나면 Docker 빌드를 다시 하여 `chaincode` 이미지를 최신 버전으로 업데이트
