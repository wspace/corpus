FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /I_C_Whitespace
COPY I_C_Whitespace .
RUN make

FROM scratch

COPY --from=builder /I_C_Whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
