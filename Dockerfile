FROM node:lts-alpine as front-end
WORKDIR /openpbl-landing
COPY openpbl-landing/package.json package.json
RUN yarn install
COPY openpbl-landing/. .
RUN set NODE_OPTIONS=--max_old_space_size=2048
RUN yarn build


FROM golang:alpine3.13
ENV RUNMODE=prod
WORKDIR /OpenPBL
RUN  go env -w GO111MODULE=on ;go env -w GOPROXY=https://goproxy.cn,direct
RUN apk add --update alpine-sdk
COPY go.mod .
RUN go mod download
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
COPY --from=front-end /openpbl-landing/build openpbl-landing/build
RUN go build
ENTRYPOINT ["./OpenPBL"]
