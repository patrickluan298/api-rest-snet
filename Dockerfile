FROM golang:1.18-alpine
WORKDIR /api
COPY . .
RUN apk add install go
RUN go run main.go
ENTRYPOINT ["postgres"]