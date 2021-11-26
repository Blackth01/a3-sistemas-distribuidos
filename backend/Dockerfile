FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . ./

RUN go build -o /backend

CMD ["/backend"]