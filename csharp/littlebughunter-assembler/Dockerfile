FROM mono AS builder

WORKDIR /WhitespaceAssembler
COPY WhitespaceAssembler .
RUN msbuild /p:Configuration=Release

FROM mono:slim

WORKDIR /WhitespaceAssembler
COPY --from=builder /WhitespaceAssembler/bin/Release/WhitespaceAssembler.exe .
ENTRYPOINT ["mono", "WhitespaceAssembler.exe"]
