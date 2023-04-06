FROM golang:1.20-alpine
ENV ENV_APP=""

WORKDIR /app
COPY . .

RUN apk update && apk add --no-cache git

RUN go mod download && go mod tidy

RUN go build -o /todo-go-rest

CMD ["/todo-go-rest"]