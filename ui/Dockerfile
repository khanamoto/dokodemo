FROM node:14.5-slim

ENV WORKDIR /usr/src/app
WORKDIR $WORKDIR

ENV PATH $WORKDIR/node_modules/.bin:$PATH

COPY package*.json ./

RUN npm install --silent
#本番用
#RUN npm install --only=production

COPY . $WORKDIR

CMD [ "npm", "start" ]