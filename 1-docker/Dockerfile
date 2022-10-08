FROM node:lts-alpine
#create the usergroup and user. Add the user to the newly created group
RUN addgroup nonroot && adduser -G nonroot nrUser 
WORKDIR /app
#copy the below files first before runing Yarn install so they can be cached
COPY ./package.json package.json
COPY ./yarn.lock yarn.lock
RUN yarn install
#index file should be copied after since it would be changed more often
COPY index.js
#switch to the user
USER username
CMD yarn start
