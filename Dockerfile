FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk update && apk upgrade -U
RUN apk --no-cache add ca-certificates libc6-compat 

WORKDIR /root/

COPY --from=builder /app/main .
COPY ./views /root/views
COPY ./public /root/public

RUN chmod +x /root/main

EXPOSE 3000

CMD ["./main --command=server"]
