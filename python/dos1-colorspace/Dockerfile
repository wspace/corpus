FROM python:2-alpine

RUN apk add gcc musl-dev jpeg-dev zlib-dev
RUN pip install pillow
WORKDIR /colorspace
COPY colorspace run.sh ./
ENTRYPOINT ["./run.sh"]
