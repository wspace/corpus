FROM python:2-alpine

RUN apk add --no-cache gcc musl-dev jpeg-dev zlib-dev
RUN pip install pillow
WORKDIR /colorspace
COPY colorspace run.sh ./
ENTRYPOINT ["./run.sh"]
