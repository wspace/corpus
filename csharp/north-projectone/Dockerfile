FROM mono AS builder

WORKDIR /WhitespaceInterpreter
COPY projectOne/The-Code/WhitespaceInterpreter .
RUN msbuild /p:Configuration=Debug

FROM scratch

COPY --from=builder /WhitespaceInterpreter/bin/Debug/WhitespaceInterpreter.exe /
ENTRYPOINT ["/WhitespaceInterpreter.exe"]
