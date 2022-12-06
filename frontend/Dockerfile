FROM node:19-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent -g npm@9.1.2 react-scripts@5.0.1 \
    && npm install --silent \
    && chown -R 744 /app/node_modules

COPY . ./

CMD ["npm", "start"]
