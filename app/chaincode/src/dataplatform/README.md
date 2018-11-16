# chaincode 

```
# AWS CLI를 통해 레지스트리에 로그인
$ aws ecr get-login --no-include-email --region us-east-1

# 소스 코드를 포함하여 빌드
$ sudo docker build -t chaincode .

# 레지스트리에 태깅 및 업로드
$ sudo docker tag chaincode:DataXchain/chaincode:latest
$ sudo docker push DataXchain/chaincode:latest
```
