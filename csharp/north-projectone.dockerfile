FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/North15/projectOne
WORKDIR /home/projectOne/The-Code/WhitespaceInterpreter
RUN msbuild /p:Configuration=Debug
RUN test -f /home/projectOne/The-Code/WhitespaceInterpreter/bin/Debug/WhitespaceInterpreter.exe
