FROM python:3.12
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
RUN adduser django
WORKDIR /code
COPY requirements.txt /code
RUN pip install -r requirements.txt