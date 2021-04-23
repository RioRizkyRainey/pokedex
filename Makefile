.PHONY: all build-protoc


PROTO_FILE=$(shell find . -type f -name "*.proto");
build-protoc: $(PROTO_FILE)
	for file in $^ ; do \
		protoc -I $(shell dirname $<) --go_out=$(shell dirname $<) --go-grpc_out=$(shell dirname $<) $$file; \
	done;
	
pokemon:
	go run cmd/pokemon/main/main.go

moves:
	go run cmd/moves/main/main.go

attack:
	go run cmd/attack/main/main.go

gateway:
	go run cmd/gateway/main/main.go

run:
	@docker-compose up -d gateway

test:
	go install github.com/newm4n/goornogo
	export GO111MODULE on; \
	go test ./... -cover -vet -all -v -short -covermode=count -coverprofile=coverage.out
	goornogo -i coverage.out -c 30