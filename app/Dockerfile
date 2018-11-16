FROM node:8.11.4-stretch

WORKDIR /usr/src/app

COPY . .


RUN cd ./web && npm install
RUN cd ./middleware && npm install