FROM alpine as builder

RUN apk add git make gcc musl-dev flex bison
RUN git clone https://gitlab.com/tejaskasetty/ws-compiler
WORKDIR /ws-compiler/flex-bison/262_267_256
RUN make

FROM scratch as runner

COPY --from=builder /ws-compiler/flex-bison/262_267_256/ws /
ENTRYPOINT ["/ws"]
