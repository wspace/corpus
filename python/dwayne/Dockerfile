FROM python:3.7-alpine

RUN pip install pipenv
WORKDIR /whitespace-python
COPY whitespace-python .
RUN pipenv install --dev
RUN python -m unittest
ENTRYPOINT ["python", "-m", "whitespace"]
