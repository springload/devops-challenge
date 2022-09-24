FROM node:lts-alpine

# Add a non-root user and switch
RUN addgroup -S nonrootgroup
RUN adduser -S -D nonrootuser nonrootgroup
USER nonrootuser

WORKDIR /app

COPY ./yarn.lock ./yarn.lock
COPY ./package.json ./package.json
RUN yarn install

COPY ./ ./
CMD yarn start
