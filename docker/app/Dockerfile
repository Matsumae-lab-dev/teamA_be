FROM golang:1.20.3-alpine

COPY app /app/
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]