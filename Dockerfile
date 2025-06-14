FROM golang:1.24-alpine

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o flashbook .

EXPOSE 8080

CMD ["./flashbook"]