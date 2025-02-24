FROM openresty/openresty:alpine	

RUN apk add perl curl \
    && apk cache clean \
    && opm get ledgetech/lua-resty-http

COPY ./nginx/conf/nginx.conf /usr/local/openresty/nginx/conf
COPY ./nginx/lua /usr/local/openresty/nginx/lua
