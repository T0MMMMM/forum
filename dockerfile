FROM golang:1.22.3-alpine AS builder

RUN apk update && apk add --no-cache gcc musl-dev

CMD mkdir /app
WORKDIR /app

COPY . .
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/gofiber/websocket/v2
RUN go get github.com/gofiber/template/html/v2
RUN go get modernc.org/sqlite

RUN go build -o main .

FROM alpine:latest
CMD mkdir /app
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/serv ./serv

EXPOSE 8080
ENTRYPOINT ["./main"]