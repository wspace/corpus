FROM node:alpine

WORKDIR /Whitespace-JS
COPY Whitespace-JS .
RUN npm install

ENTRYPOINT ["node", "bin/wspace"]
# ENTRYPOINT ["node", "bin/unws2ws"]
# ENTRYPOINT ["node", "bin/ws2unws"]
