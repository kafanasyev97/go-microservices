PROTO_REPO = github.com/kafanasyev97/go-microservices-proto
PROTO_PATH = ./proto
AUTH_PROTO = auth/auth.proto

generate:
	rm -rf $(PROTO_PATH)/auth
	mkdir -p $(PROTO_PATH)
	git clone https://$(PROTO_REPO) $(PROTO_PATH)/tmp
	protoc -I $(PROTO_PATH)/tmp \
		--go_out=$(PROTO_PATH) \
		--go-grpc_out=$(PROTO_PATH) \
		$(PROTO_PATH)/tmp/$(AUTH_PROTO)
	rm -rf $(PROTO_PATH)/tmp
