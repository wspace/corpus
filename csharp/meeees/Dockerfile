FROM mono AS builder

WORKDIR /Whitespace
COPY Whitespace .
RUN msbuild /p:Configuration=Debug

FROM scratch

COPY --from=builder /Whitespace/bin/Debug/Whitespace.exe /
ENTRYPOINT ["/Whitespace.exe"]
