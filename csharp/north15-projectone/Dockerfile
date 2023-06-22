FROM mono AS builder

WORKDIR /WhitespaceInterpreter
COPY projectOne/The-Code/WhitespaceInterpreter .
RUN msbuild /p:Configuration=Release

FROM mono:slim

WORKDIR /WhitespaceInterpreter
COPY --from=builder /WhitespaceInterpreter/bin/Release/WhitespaceInterpreter.exe .
ENTRYPOINT ["mono", "WhitespaceInterpreter.exe"]
