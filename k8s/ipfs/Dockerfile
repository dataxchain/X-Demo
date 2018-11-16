# https://github.com/yeasy/docker-ipfs
#
# Dockerfile for IPFS (https://ipfs.io/)
# Data is stored under /root/.ipfs/


FROM golang:1.10
#LABEL maintainer "Baohua Yang <yangbaohua@gmail.com>"

#ENV DEBIAN_FRONTEND noninteractive

ENV SWARM_PORT 4001
ENV API_PORT 5001
ENV GATEWAY_PORT 8080

EXPOSE ${SWARM_PORT}
EXPOSE ${API_PORT}
EXPOSE ${GATEWAY_PORT}

#ENV IPFS_PATH /root/.ipfs
#VOLUME $IPFS_PATH

# Install ipfs using ipfs-update
RUN go get -u github.com/ipfs/ipfs-update \
    && ipfs-update install latest

COPY docker-entrypoint.sh /usr/local/bin/
#RUN ln -s /usr/local/bin/docker-entrypoint.sh /
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["ipfs", "daemon"]
