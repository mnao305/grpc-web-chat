FROM node:14.15.3-alpine3.10

WORKDIR /client

COPY . .

RUN npm install -g http-server

RUN npm ci


RUN npm run build

CMD [ "http-server", "dist" ]