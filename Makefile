PROTO_DIR := proto
OUT_DIR := .

PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

.PHONY: proto

proto:
	protoc \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)