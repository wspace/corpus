FROM mono AS builder

WORKDIR /Whitespace-Interpreter
COPY Whitespace-Interpreter .
RUN mcs -debug -out:WhitespaceInterpreter.exe WhitespaceInterpreter/*.cs

FROM scratch

COPY --from=builder /Whitespace-Interpreter/WhitespaceInterpreter.exe /
ENTRYPOINT ["/WhitespaceInterpreter.exe"]
