FROM redis/redis-stack

WORKDIR /redis

#COPY redis.conf /usr/local/etc/redis/redis.conf

COPY init.sh ./

RUN chmod +x init.sh
