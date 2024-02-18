FROM alpine

RUN apk add --no-cache lua5.4
WORKDIR /WhitespaceInterpreter
COPY WhitespaceInterpreter .
ENTRYPOINT ["lua5.4", "main.lua"]
