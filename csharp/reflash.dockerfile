FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/reflash-blog/WhiteSpaceInterpreter
WORKDIR /home/WhiteSpaceInterpreter/WhiteSpaceInterpretator
RUN msbuild /p:Configuration=Debug
RUN test -f /home/WhiteSpaceInterpreter/WhiteSpaceInterpretator/bin/Debug/WhiteSpaceInterpretator.exe
