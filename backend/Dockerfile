FROM golang:1.23.0 AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM debian:stable-slim AS app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
EXPOSE 8080
COPY --from=build /go/src/app/app .
COPY --from=build /go/src/app/FirebaseAdminSDK.json .
CMD ["./app"]