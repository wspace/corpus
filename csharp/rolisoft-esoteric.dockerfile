FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/RoliSoft/Esoteric-Code-Interpreter
WORKDIR /home/Esoteric-Code-Interpreter
RUN msbuild /p:Configuration=Debug\;AssemblyName=EsotericCodeInterpreter
RUN test -f /home/Esoteric-Code-Interpreter/bin/Debug/EsotericCodeInterpreter.exe
