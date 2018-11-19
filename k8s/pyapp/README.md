### 1. 개발

```
$ cd k8s
$ kubectl create -f pyapp.yaml

$ telepresence --swap-deployment pyapp --expose 31641:5000 --docker-run --rm -it \
-v ../../dataplatform/app:/usr/src/app \
DataXchain/dataplatform:latest bash

# telepresence를 통해 다음과 같이 배포 환경과 동일한 환경에서 개발 중인 모듈 실행 가능
root@4a7c97ee5243:/usr/src/app# python web.py
```

- `dataplatform/app` 디렉토리가 컨테이너의 `/usr/src/app`과 연결되어있으므로 코드 수정시 컨테이너에 바로 반영됨
- 리눅스 서버의 `dataplatform/app` 을 sftp로 Pycharm과 연동하면 편리하게 사용할 수 있음

### 2. 배포

- 개발이 끝나면 Docker 빌드를 다시 하여 `dataplatform` 이미지를 최신 버전으로 업데이트
