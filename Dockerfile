FROM alpine:3.15
RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release
WORKDIR /app
COPY defaultProject defaultProject
COPY config/* config/.

LABEL Name=defaultProject Version=0.0.1
ENTRYPOINT ["/app/defaultProject"]