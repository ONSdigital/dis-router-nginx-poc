FROM node:22-alpine

RUN apk add --no-cache nginx

WORKDIR /app

COPY /sidecar/package.json /sidecar/package-lock.json ./
RUN npm install

COPY /sidecar/index.js .

COPY ../nginx/nginx.conf /etc/nginx/nginx.conf
COPY ../nginx/redirects.conf /etc/nginx/redirects.conf
COPY ../nginx/routes.conf /etc/nginx/routes.conf

EXPOSE 80 5000

CMD ["sh", "-c", "nginx && node index.js"]
