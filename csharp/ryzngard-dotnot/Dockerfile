FROM mono AS builder

WORKDIR /DotNot
COPY DotNot .
RUN nuget restore
WORKDIR /DotNot/src/Whitespace.TestProject
RUN msbuild /p:Configuration=Release
