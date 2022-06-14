FROM alpine

RUN apk add git openjdk8 erlang
RUN git clone https://github.com/wspace/bearice-grassmudhorse grass-mud-horse
WORKDIR /grass-mud-horse
RUN javac -d bin src/cn/icybear/GrassMudHorse/*.java
RUN erlc erlang/grass_mud_horse.erl

COPY bearice-grassmudhorse.sh run.sh
ENTRYPOINT ["./run.sh"]
