FROM node:alpine

WORKDIR /whitespace-lang
COPY whitespace-lang .
RUN npm install

ENTRYPOINT ["node", "index.js"]
