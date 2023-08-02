FROM eclipse-temurin:20-jdk-alpine AS builder

WORKDIR /grass-mud-horse
COPY grass-mud-horse .
RUN javac -d bin src/cn/icybear/GrassMudHorse/*.java && \
    jar -cfe GrassMudHorse.jar cn.icybear.GrassMudHorse.AlpacaVM -C bin .

FROM eclipse-temurin:20-jre-alpine

RUN apk add --no-cache erlang
WORKDIR /grass-mud-horse
COPY --from=builder /grass-mud-horse/GrassMudHorse.jar /grass-mud-horse/erlang ./
COPY run.sh grass-mud-horse/*.gmh grass-mud-horse/*.ws ./
RUN erlc grass_mud_horse.erl && \
    rm grass_mud_horse.erl grass_mud_horse.hrl
ENTRYPOINT ["./run.sh"]
