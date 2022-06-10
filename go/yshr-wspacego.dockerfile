FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/135yshr/wspacego
WORKDIR /home/wspacego
RUN go mod init github.com/135yshr/wspacego
RUN go mod tidy
RUN go build
# builds: /home/wspacego/wspacego
