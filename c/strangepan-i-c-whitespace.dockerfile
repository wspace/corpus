FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/StrangePan/I_C_Whitespace
WORKDIR /home/I_C_Whitespace
RUN make
RUN test -f /home/I_C_Whitespace/whitespace
