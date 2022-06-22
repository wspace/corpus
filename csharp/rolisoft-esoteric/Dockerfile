FROM mono AS builder

WORKDIR /Esoteric-Code-Interpreter
COPY Esoteric-Code-Interpreter .
RUN msbuild /p:Configuration=Release\;AssemblyName=EsotericCodeInterpreter

FROM mono:slim

WORKDIR /Esoteric-Code-Interpreter
COPY --from=builder /Esoteric-Code-Interpreter/bin/Release/EsotericCodeInterpreter.exe .
ENTRYPOINT ["mono", "EsotericCodeInterpreter.exe"]
