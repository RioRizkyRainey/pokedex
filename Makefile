.PHONY: all build-protoc


PROTO_FILE=$(shell find . -type f -name "*.proto");
build-protoc: $(PROTO_FILE)
	for file in $^ ; do \
		protoc -I $(shell dirname $<) --go_out=$(shell dirname $<) --go-grpc_out=$(shell dirname $<) $$file; \
	done;
	