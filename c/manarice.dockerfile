FROM alpine

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/ManaRice/limelib /home/utils/limelib
WORKDIR /home/utils/limelib
RUN make
RUN git clone https://github.com/ManaRice/whitespace /home/utils/whitespace
WORKDIR /home/utils/whitespace
RUN make
RUN test -f /home/utils/whitespace/lwsvm
RUN test -f /home/utils/whitespace/lwsa
