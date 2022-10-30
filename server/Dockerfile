FROM golang

WORKDIR /usr/app/
COPY ./ ./
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "static"' -o server-mirror

FROM alpine
RUN apk upgrade --update-cache --available && \
    apk add openssl && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=0 /usr/app/server-mirror ./
EXPOSE 8080
ENTRYPOINT ["/app/server-app"]