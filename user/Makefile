include .env

LOCAL_BIN:=$(CURDIR)/bin
DOCKER_IMAGE_NAME=user
DOCKERFILE_PATH=Dockerfile
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_PG_USER=$(POSTGRES_USER)
LOCAL_PG_PASSWORD=$(POSTGRES_PASSWORD)
LOCAL_PG_DB=$(POSTGRES_DB)
LOCAL_PG_PORT=$(PG_PORT)
LOCAL_MIGRATION_DSN=postgres://$(LOCAL_PG_USER):$(LOCAL_PG_PASSWORD)@localhost:$(LOCAL_PG_PORT)/$(LOCAL_PG_DB)?sslmode=disable

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

build-docker:
	docker build -t $(DOCKER_IMAGE_NAME) -f $(DOCKERFILE_PATH) .

run-docker:
	docker-compose up -d

stop-docker:
	docker-compose down

docker: build-docker run-docker

clean-docker: stop-docker
