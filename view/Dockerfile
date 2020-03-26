# vue,vuetify開発環境
FROM node:10.12-alpine as build-stage

WORKDIR /app
COPY . .

ENV https_proxy "http://sg-sh128-1.isid.co.jp:8080"
ENV http_proxy "http://sg-sh128-1.isid.co.jp:8080"

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

CMD ["/bin/ash"]