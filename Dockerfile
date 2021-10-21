FROM alpine:3.14
RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release
WORKDIR /app
COPY <project-name> <project-name>
COPY config.yaml config.yaml

LABEL Name=<project-name> Version=0.0.1
ENTRYPOINT ["/app/<project-name>"]