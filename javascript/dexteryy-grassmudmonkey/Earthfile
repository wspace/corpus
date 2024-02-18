FROM node:alpine

WORKDIR /GrassMudMonkey
COPY GrassMudMonkey cli.js ./
RUN cat cli.js >> grassmudmonkey.js

ENTRYPOINT ["node", "grassmudmonkey"]
