migrate-file:
	migrate create -ext sql -dir migrations ${filename}

swagger:
	swag init --parseDependency --parseInternal

install:
	sh install.sh

run:
	go run main.go start --config=./config/local/config.yaml

wire:
	wire ./internal

compose:
	docker compose -f docker-compose.local.yaml up -d

compose-dev:
	docker compose -f docker-compose.dev.yaml up -d

mock:
	mockgen --build_flags=--mod=mod --destination=./internal/domain/interfaces/common/mocks/mock.go github.com/gogovan/ggx-kr-payment-service/internal/domain/interfaces/fileservice FileService

test:
	go test -v ./...

build-image:
	docker build . -t common-service

proto-gen:
	protoc --go_out=pkg/grpc  --go_opt=paths=source_relative \
        --go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
        -I pkg/grpc/proto \
        pkg/grpc/proto/*.proto
