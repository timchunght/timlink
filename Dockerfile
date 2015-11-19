# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/timchunght/timlink

# Build the golang-docker command inside the container.
RUN go install github.com/timchunght/timlink

EXPOSE 8080
# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/timlink

# http server listens on port 8080.
