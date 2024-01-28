FROM golang:1.21.6-bullseye

RUN curl -SsfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# add user
RUN useradd -m -s /bin/bash -u 1000 app
USER app

WORKDIR /app

COPY ./source /app
RUN go mod download
RUN go build -o /app/fiber /app/main.go