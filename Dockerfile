# Use a multi-stage build
FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o ./todo-api .

FROM gcr.io/distroless/static:nonroot

# Copy exe from build container
COPY --from=builder /app/todo-api ./

# Define start command
CMD ["./todo-api"]
