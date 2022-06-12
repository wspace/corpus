FROM mono as builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/DenisLabrecque/Whitespace-Interpreter
WORKDIR /Whitespace-Interpreter
RUN mcs -debug -out:WhitespaceInterpreter.exe WhitespaceInterpreter/*.cs

FROM scratch as runner

COPY --from=builder /Whitespace-Interpreter/WhitespaceInterpreter.exe /
ENTRYPOINT ["/WhitespaceInterpreter.exe"]
