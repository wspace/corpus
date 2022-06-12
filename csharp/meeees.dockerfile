FROM mono as builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/wspace/meeees-csharp Whitespace
WORKDIR /Whitespace
RUN msbuild /p:Configuration=Debug

FROM scratch as runner

COPY --from=builder /Whitespace/bin/Debug/Whitespace.exe /
ENTRYPOINT ["/Whitespace.exe"]
