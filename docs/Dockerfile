FROM python:3.8 AS builder

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

FROM builder
COPY mkdocs.yml mkdocs.yml
COPY files files

EXPOSE 8000
CMD [ "mkdocs", "serve", "-a0.0.0.0:8000" ]