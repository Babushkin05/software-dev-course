PROTO_DIR=api
OUT_DIR=gen

proto:
	protoc \
		-I $(PROTO_DIR) \
		--go_out=$(PROTO_DIR)/$(OUT_DIR) \
		--go-grpc_out=$(PROTO_DIR)/$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/orders/orders.proto \
		$(PROTO_DIR)/payments/payments.proto

clean:
	rm -rf $(PROTO_DIR)/$(OUT_DIR)/*