FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/littleBugHunter/WhitespaceAssembler
WORKDIR /home/WhitespaceAssembler
RUN msbuild /p:Configuration=Debug
RUN test -f /home/WhitespaceAssembler/bin/Debug/WhitespaceAssembler.exe
