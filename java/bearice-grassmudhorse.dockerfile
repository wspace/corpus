FROM alpine

RUN apk add git openjdk8 erlang
RUN git clone https://github.com/wspace/bearice-grassmudhorse grass-mud-horse

WORKDIR /grass-mud-horse
RUN javac -d bin src/cn/icybear/GrassMudHorse/*.java

WORKDIR /grass-mud-horse/erlang
RUN erlc grass_mud_horse.erl
RUN echo '#!/bin/sh' > run.sh
RUN echo 'exec /usr/bin/erl -noshell -s grass_mud_horse compile_run "$1" -s init stop' >> run.sh
RUN chmod +x run.sh

WORKDIR /grass-mud-horse
ENTRYPOINT ["/usr/bin/java", "-cp", "bin", "cn.icybear.GrassMudHorse.AlpacaVM"]
# ENTRYPOINT ["/usr/bin/java", "-cp", "bin", "cn.icybear.GrassMudHorse.JOTCompiler"]
# ENTRYPOINT ["/usr/bin/java", "-cp", "bin", "cn.icybear.GrassMudHorse.WS2GMH"]
# ENTRYPOINT ["./run.sh"]
