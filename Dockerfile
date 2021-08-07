FROM golang:1.16-rc-buster
WORKDIR /openpbl
RUN go env -w CGO_ENABLED=0 GOPROXY=https://goproxy.io,direct GOOS=linux GOARCH=amd64 \
    && apt update && apt install sudo \
    && wget https://nodejs.org/dist/v12.22.0/node-v12.22.0-linux-x64.tar.gz \
    && sudo tar xf node-v12.22.0-linux-x64.tar.gz \
    && sudo apt install wait-for-it
ENV PATH=$PATH:/openpbl/node-v12.22.0-linux-x64/bin
RUN npm install -g yarn

COPY openpbl-landing/package.json /openpbl/openpbl-landing/package.json
RUN cd openpbl-landing && yarn install

COPY openpbl-landing /openpbl/openpbl-landing
RUN cd openpbl-landing && yarn build && rm -rf node_modules

COPY ./ /openpbl
RUN cd /openpbl && go build main.go


FROM alpine:3.7
COPY --from=0 /openpbl   /
COPY --from=0 /usr/bin/wait-for-it  /
RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add bash
ENV RUNMODE=prod
CMD ./wait-for-it openpbl-db:3308 && ./main