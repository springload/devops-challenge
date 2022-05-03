FROM companyrepo/app_yarn_dependencies:latest 

WORKDIR /app

COPY ./ ./

RUN chown -R node:node /app

USER node
RUN yarn install

CMD yarn start
