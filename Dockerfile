# Development Dockerfile with Air for live reload
FROM golang:1.25-alpine AS development

# Install Air for live reload
RUN go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Expose port (adjust as needed)
EXPOSE 8080

# Run Air for live reload
CMD ["air", "-c", ".air.toml"]
