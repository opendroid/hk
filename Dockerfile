# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Sample https://cloud.google.com/run/docs/quickstarts/build-and-deploy?_ga=2.91290522.-1679093051.1593441137

# Start from the latest golang base image
FROM golang:1.19 as builder
LABEL maintainer="Open Web <plutoapps@outlook.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go test ./... && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o hk

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/hk /hk
COPY --from=builder /app/public /public
EXPOSE 8080
CMD ["/hk", "-port=8080"]
