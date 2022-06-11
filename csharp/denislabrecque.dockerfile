FROM mono

RUN apt-get update
RUN apt-get install -y git
WORKDIR /home
RUN git clone https://github.com/DenisLabrecque/Whitespace-Interpreter
WORKDIR /home/Whitespace-Interpreter
RUN mcs -debug -out:WhitespaceInterpreter.exe WhitespaceInterpreter/*.cs
RUN test -f /home/Whitespace-Interpreter/WhitespaceInterpreter.exe
