FROM node:19.5.0-slim
WORKDIR /usr/app
COPY package.json .
RUN yarn install --quiet
COPY . .