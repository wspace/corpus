FROM python:2.7-alpine

WORKDIR /abyad
COPY abyad .
RUN python generate_programs
ENTRYPOINT ["python", "-m", "src"]
