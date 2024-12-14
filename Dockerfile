# Build the application from source
FROM golang:1.23 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Generate swagger docs
RUN swag init -g cmd/main.go -o internal/docs
RUN mv ./internal/docs/swagger.yaml ./internal/docs/openapi.yaml
RUN mv ./internal/docs/swagger.json ./internal/docs/openapi.json

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/main.go

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
WORKDIR /

COPY --from=build-stage /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]