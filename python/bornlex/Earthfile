FROM python:3-alpine

WORKDIR /Whitespace-interpreter
COPY Whitespace-interpreter .
RUN python tests.py
ENTRYPOINT ["python", "main.py"]
