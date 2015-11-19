# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/timlink
WORKDIR /go/src/timlink
# Build the golang-docker command inside the container.
RUN go install timlink

# Run the golang-docker command when the container starts.

# http server listens on port 8080.
