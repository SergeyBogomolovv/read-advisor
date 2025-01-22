.PHONY: gen-proto
gen-proto:
	@name=$(name);
	mkdir -p lib/api/gen/$(name);
	@protoc --proto_path=lib/api/proto \
	  --go_out=lib/api/gen/$(name) --go_opt=paths=source_relative \
		--go-grpc_out=lib/api/gen/$(name) --go-grpc_opt=paths=source_relative \
		lib/api/proto/$(name).proto