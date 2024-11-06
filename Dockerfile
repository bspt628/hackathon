FROM golang:1.22.3 AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GDOS=linux go build -o app .

FROM debian:stable-slim AS app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
EXPOSE 9000
COPY --from=build /go/src/app/app .
CMD ["./app"]