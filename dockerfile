FROM golang:1.16
WORKDIR /openpbl
COPY ./ /openpbl
RUN go env -w CGO_ENABLED=0 GOPROXY=https://goproxy.io,direct GOOS=linux GOARCH=amd64 \
    && apt update && apt install sudo \
    && wget https://nodejs.org/dist/v12.22.0/node-v12.22.0-linux-x64.tar.gz \
    && sudo tar xf node-v12.22.0-linux-x64.tar.gz \
    && sudo apt install wait-for-it
ENV PATH=$PATH:/openpbl/node-v12.22.0-linux-x64/bin
RUN npm install -g yarn \
    && cd openpbl-landing \
    && yarn install \
    && yarn run build \
    && rm -rf node_modules \
    && cd /openpbl \
    && go build main.go


FROM alpine:3.7
ENV RUNMODE=prod
COPY --from=0 /openpbl   /
COPY --from=0 /usr/bin/wait-for-it  /
RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add bash
CMD ./wait-for-it db:3306 && ./main