FROM python:3-alpine

RUN apk add --no-cache make g++ gmp-dev
RUN pip install -v z3-solver
WORKDIR /whitesymex
COPY whitesymex .
ENTRYPOINT ["python", "-m", "whitesymex.cli"]
