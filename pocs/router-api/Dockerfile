FROM node:22

RUN npm install -g json-server

COPY ./run.sh /run.sh

ENTRYPOINT ["bash", "./run.sh"]
