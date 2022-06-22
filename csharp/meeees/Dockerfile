FROM mono AS builder

WORKDIR /Whitespace
COPY Whitespace .
RUN msbuild /p:Configuration=Release

FROM mono:slim

WORKDIR /Whitespace
COPY --from=builder /Whitespace/bin/Release/Whitespace.exe .
ENTRYPOINT ["mono", "Whitespace.exe"]
