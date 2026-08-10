# Use a smaller base image for the build stage
FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY . .

# Enable Go modules
ENV GO111MODULE=on
RUN go mod tidy
# Build the Go app and strip debugging information to reduce size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/action ./

# Distroless runtime image (statically linked binary, no shell/libc bloat)
FROM gcr.io/distroless/static-debian12:nonroot

# Copy the compiled Go program from the builder stage
COPY --from=builder /bin/action /app/action

COPY --from=builder /src/example.yaml /example.yaml

# Set entrypoint
ENTRYPOINT ["/app/action", "run"]
