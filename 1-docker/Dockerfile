FROM node:lts-alpine

WORKDIR /app

COPY ./yarn.lock ./yarn.lock
COPY ./package.json ./package.json
RUN yarn install

COPY ./ ./
CMD yarn start
