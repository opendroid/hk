# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Sample https://cloud.google.com/run/docs/quickstarts/build-and-deploy?_ga=2.91290522.-1679093051.1593441137

# Start from the latest golang base image
FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="Open Web <plutoapps@outlook.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and  go.sum files
COPY go.mod ./
RUN go mod download
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o hk

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/hk /hk
COPY --from=builder /app/public /public

# Expose port 8888 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/hk", "-file=/public/sampleexport.xml"]
