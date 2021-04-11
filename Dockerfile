# Our server is written in Go so we will use the Go base image. All 
# Docker images start from a base image.
FROM golang:1.16

# Sets all future commands to work relative to /app.
WORKDIR /app

# Copy over the files from this directory (the first dot)
# to the app directory (the second dot).
ADD . .

# Download all the dependencies for the server.
RUN go mod download

# Build the Go server so we can run it. Save it as main in the 
# current directory (/app).
RUN go build -o main .

# Let Docker know our server is listening on port 80 internally.
EXPOSE 80

# When the container starts, this starts the server.
CMD ["./main"]
