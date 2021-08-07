FROM golang:1.16.7-alpine3.14 AS step01
WORKDIR /demo
COPY . .
RUN go mod tidy  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/main.go

FROM alpine:3.14.1  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=step01 /demo/api ./api
EXPOSE 8080
CMD ["./api"]