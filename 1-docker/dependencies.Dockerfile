FROM node:lts-alpine

WORKDIR /app

COPY ./ ./

RUN chown -R node:node /app

USER node
RUN yarn install

USER root
CMD /bin/true
