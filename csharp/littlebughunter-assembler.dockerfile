FROM mono as builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/littleBugHunter/WhitespaceAssembler
WORKDIR /WhitespaceAssembler
RUN msbuild /p:Configuration=Debug

FROM scratch as runner

COPY --from=builder /WhitespaceAssembler/bin/Debug/WhitespaceAssembler.exe /
ENTRYPOINT ["/WhitespaceAssembler.exe"]
