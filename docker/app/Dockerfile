FROM golang:1.20.3-alpine

# COPY app /app/
# WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest

# CMD ["air"]

COPY app /app/
WORKDIR /app

RUN go mod download && go mod verify

CMD ["go", "run", "main.go"] 