FROM mono AS builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/North15/projectOne
WORKDIR /projectOne/The-Code/WhitespaceInterpreter
RUN msbuild /p:Configuration=Debug

FROM scratch

COPY --from=builder /projectOne/The-Code/WhitespaceInterpreter/bin/Debug/WhitespaceInterpreter.exe /
ENTRYPOINT ["/WhitespaceInterpreter.exe"]
