FROM mono as builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/reflash-blog/WhiteSpaceInterpreter
WORKDIR /WhiteSpaceInterpreter/WhiteSpaceInterpretator
RUN msbuild /p:Configuration=Debug

FROM scratch as runner

COPY --from=builder /WhiteSpaceInterpreter/WhiteSpaceInterpretator/bin/Debug/WhiteSpaceInterpretator.exe /
ENTRYPOINT ["/WhiteSpaceInterpretator.exe"]
