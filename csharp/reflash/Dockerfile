FROM mono AS builder

WORKDIR /WhiteSpaceInterpreter
COPY WhiteSpaceInterpreter .
WORKDIR /WhiteSpaceInterpreter/WhiteSpaceInterpretator
RUN msbuild /p:Configuration=Release

FROM mono:slim

WORKDIR /WhiteSpaceInterpreter
COPY --from=builder /WhiteSpaceInterpreter/WhiteSpaceInterpretator/bin/Release/WhiteSpaceInterpretator.exe .
ENTRYPOINT ["mono", "WhiteSpaceInterpretator.exe"]
