FROM mono AS builder

WORKDIR /Esoteric-Code-Interpreter
COPY Esoteric-Code-Interpreter .
RUN msbuild /p:Configuration=Debug\;AssemblyName=EsotericCodeInterpreter

FROM scratch

COPY --from=builder /Esoteric-Code-Interpreter/bin/Debug/EsotericCodeInterpreter.exe /
ENTRYPOINT ["/EsotericCodeInterpreter.exe"]
