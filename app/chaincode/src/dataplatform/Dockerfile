FROM golang:1.10.3-stretch

WORKDIR /go/src/app

RUN go get github.com/hyperledger/fabric/core/chaincode/shim \
        && go get github.com/hyperledger/fabric/protos/peer

COPY . .