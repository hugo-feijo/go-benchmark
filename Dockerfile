FROM golang:1.16.5 as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY *.go ./
# Expose port
EXPOSE 4000
# Start app
CMD go run .
