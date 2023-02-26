FROM alpine

RUN apk add --no-cache openjdk8 erlang
WORKDIR /grass-mud-horse
COPY grass-mud-horse run.sh ./
RUN javac -d bin src/cn/icybear/GrassMudHorse/*.java
RUN erlc erlang/grass_mud_horse.erl

ENTRYPOINT ["./run.sh"]
