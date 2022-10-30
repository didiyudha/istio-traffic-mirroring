server-app: 
	cd server; GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "static"' -o $@
client-app:
	cd client; GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "static"' -o $@

.PHONY: client.image
client.image:
	docker build -t didiyudha/client-app:v1.0.0 ./client

.PHONY: server.image
server.image:
	docker build -t didiyudha/server-app:v1.0.0 ./server

.PHONY: server.mirror.image
server.mirror.image:
	docker build -t didiyudha/server-app:mirror ./server