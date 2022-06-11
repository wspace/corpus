FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/wspace/meeees-csharp Whitespace
WORKDIR /home/Whitespace
RUN msbuild /p:Configuration=Debug
RUN test -f /home/Whitespace/bin/Debug/Whitespace.exe
