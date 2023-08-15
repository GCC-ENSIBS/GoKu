# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags="-w -s" -o /goku

# Expose
EXPOSE 8080

# Run
CMD ["/goku"]