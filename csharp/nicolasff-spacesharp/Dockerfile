FROM mono AS builder

RUN apt-get update && \
    apt-get install -y make flex bison
WORKDIR /spacesharp
COPY spacesharp .
RUN make MCS=mcs

FROM mono:slim

WORKDIR /spacesharp
COPY --from=builder /spacesharp/wsc.exe .
ENTRYPOINT ["mono", "wsc.exe"]
