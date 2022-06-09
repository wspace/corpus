FROM golang:1.18

RUN git clone https://github.com/135yshr/wspacego /yshr-wspacego
WORKDIR /yshr-wspacego
RUN go mod init github.com/135yshr/wspacego
RUN go mod tidy
RUN go build
# builds: /yshr-wspacego/wspacego
