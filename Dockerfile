FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/config ./config
EXPOSE 8080
ENTRYPOINT ["./main"]
