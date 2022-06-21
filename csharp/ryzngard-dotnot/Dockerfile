FROM mono AS builder

RUN apt-get update
RUN apt-get install -y git
RUN git clone https://github.com/ryzngard/DotNot
WORKDIR /DotNot
RUN nuget restore
WORKDIR /DotNot/src/Whitespace.TestProject
RUN msbuild /p:Configuration=Debug
