FROM mono AS builder

WORKDIR /WhitespaceAssembler
COPY WhitespaceAssembler .
RUN msbuild /p:Configuration=Debug

FROM scratch

COPY --from=builder /WhitespaceAssembler/bin/Debug/WhitespaceAssembler.exe /
ENTRYPOINT ["/WhitespaceAssembler.exe"]
