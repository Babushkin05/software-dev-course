SERVICE_NAME = orders-service
PROTO_DIR = api
PROTO_OUT = api/gen
PROTO_FILE = $(PROTO_DIR)/orders.proto
GO_PACKAGE = github.com/Babushkin05/software-dev-course/kr3/$(SERVICE_NAME)
DOCKER_IMAGE = $(SERVICE_NAME):latest

# Запуск приложения локально
run:
	go run ./cmd/main.go --config=config/local.yaml

# Генерация protobuf (Go + gRPC)
proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_OUT) \
		--go-grpc_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--go_opt=M$(PROTO_FILE)=$(GO_PACKAGE)/$(PROTO_OUT) \
		$(PROTO_FILE)

# Сборка Docker-образа
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Запуск контейнера
docker-run:
	docker run -p 50052:50052 --env-file .env --rm $(DOCKER_IMAGE)

# Очистка
clean:
	rm -rf $(PROTO_OUT)/*
