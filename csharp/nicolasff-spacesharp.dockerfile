FROM mono

RUN apt-get update
RUN apt-get install -y git make flex bison
WORKDIR /home
RUN git clone https://github.com/nicolasff/spacesharp
WORKDIR /home/spacesharp
RUN make MCS=mcs
RUN test -f /home/spacesharp/wsc.exe
