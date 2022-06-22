FROM python:3.7-alpine AS builder

WORKDIR /whitepy
COPY whitepy .
RUN pip install -r requirements.txt
ENTRYPOINT ["python", "whitepycli"]
