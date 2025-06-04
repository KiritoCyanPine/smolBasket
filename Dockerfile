# Use the official Go image as the base image
# Use the "alpine" variant for a lightweight image
FROM golang:1.24-alpine as builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o smolBasket ./cmd/server
RUN go build -o smol-cli ./cmd/client

# Use a minimal base image for the final container
FROM alpine:latest
WORKDIR /app
ENV PATH="/app:$PATH"
COPY --from=builder /app/smolBasket .
COPY --from=builder /app/smol-cli .
ENV TCP_PORT=9000
EXPOSE ${TCP_PORT}
CMD ["./smolBasket"]