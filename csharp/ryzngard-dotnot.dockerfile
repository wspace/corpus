FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/ryzngard/DotNot
WORKDIR /home/DotNot
RUN nuget restore
WORKDIR /home/DotNot/src/Whitespace.TestProject
RUN msbuild /p:Configuration=Debug
