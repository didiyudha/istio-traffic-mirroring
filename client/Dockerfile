FROM golang

WORKDIR /usr/app/
COPY ./ ./
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "static"' -o client-app

FROM alpine
RUN apk upgrade --update-cache --available && \
    apk add openssl && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=0 /usr/app/client-app ./
EXPOSE 3000
ENTRYPOINT ["/app/client-app"]