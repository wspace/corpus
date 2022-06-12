FROM mono as builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/RoliSoft/Esoteric-Code-Interpreter
WORKDIR /Esoteric-Code-Interpreter
RUN msbuild /p:Configuration=Debug\;AssemblyName=EsotericCodeInterpreter

FROM scratch as runner

COPY --from=builder /Esoteric-Code-Interpreter/bin/Debug/EsotericCodeInterpreter.exe /
ENTRYPOINT ["/EsotericCodeInterpreter.exe"]
