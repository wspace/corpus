FROM alpine

RUN apk add git bash g++
WORKDIR /home
RUN git clone https://github.com/ShadowMitia/whitespace
WORKDIR /home/whitespace
RUN bash build.sh release wspace
RUN test -f /home/whitespace/wspace
