FROM node:lts-alpine

# Create a new group named waseemgroup
RUN addgroup waseemgroup

# Create a new user named waseemuser and assign it to the group
RUN adduser -D -g waseemgroup waseemuser

WORKDIR /app

# Change working directory ownership  to waseemuser
RUN chown -R waseemuser:waseemgroup /app

# Change to the user "waseem"
USER waseemuser

# Abstract install packages so they can be cached
COPY package.json yarn.lock ./

COPY ./ ./

RUN yarn install

CMD yarn start
