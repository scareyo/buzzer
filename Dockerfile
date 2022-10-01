FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go mod download && \
    go build ./cmd/buzzer

FROM ubuntu:latest AS run
WORKDIR /app
COPY --from=build /app/buzzer .
CMD ["./buzzer"]
