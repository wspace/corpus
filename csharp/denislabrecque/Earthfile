FROM mono AS builder

WORKDIR /Whitespace-Interpreter
COPY Whitespace-Interpreter .
RUN mcs -optimize -out:WhitespaceInterpreter.exe WhitespaceInterpreter/*.cs

FROM mono:slim

WORKDIR /Whitespace-Interpreter
COPY --from=builder /Whitespace-Interpreter/WhitespaceInterpreter.exe .
ENTRYPOINT ["mono", "WhitespaceInterpreter.exe"]
