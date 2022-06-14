FROM mono AS builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/littleBugHunter/WhitespaceAssembler
WORKDIR /WhitespaceAssembler
RUN msbuild /p:Configuration=Debug

FROM scratch

COPY --from=builder /WhitespaceAssembler/bin/Debug/WhitespaceAssembler.exe /
ENTRYPOINT ["/WhitespaceAssembler.exe"]
