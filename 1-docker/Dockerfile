FROM node:lts-alpine
#create the usergroup and user. Add the user to the newly created group
USER root
RUN addgroup nonroot
RUN adduser -G nonroot nrUser --disabled-password
WORKDIR /app
#copy the below files first before runing Yarn install so they can be cached
COPY ./package.json package.json
COPY ./yarn.lock yarn.lock
RUN yarn install
#index file should be copied after since it would be changed more often
COPY ./index.js index.js
#switch to the user
USER nrUser
CMD yarn start