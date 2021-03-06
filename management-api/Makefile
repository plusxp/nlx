ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

API_FILES=api/management.pb.go api/management.pb.gw.go api/management.swagger.json api/mock/mock_management.go api/external/external.pb.go api/external/mock/mock_external.go
PKG_FILES=pkg/database/mock/mock_database.go pkg/directory/mock/mock_client.go pkg/management/mock/mock_client.go pkg/auditlog/mock/mock_auditlog.go pkg/txlogdb/mock/mock_database.go

.PHONY: all
all: generate migrations

.PHONY: migrations
migrations: db/migrations.go

.PHONY: generate
generate: ${API_FILES} ${PKG_FILES}

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f ${API_FILES} ${PKG_FILES} db/migrations.go

.PHONY: dependencies
dependencies:
	GO111MODULE=off go get \
		github.com/gogo/protobuf/gogoproto \
		github.com/gogo/protobuf/protoc-gen-gogofast \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

api/management.pb.go api/management.pb.gw.go api/management.swagger.json: api/management.proto
	@# Generate go, gateway
	echo $(<D)
	cd $(<D) && protoc \
		-I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:. \
		--grpc-gateway_out=\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
logtostderr=true:.\
		--swagger_out=logtostderr=true:. \
		$(<F)


api/mock/mock_management.go: api/management.pb.go
	@# Generate mock file
	mockgen -source $< -destination $@

api/external/external.pb.go: api/external/external.proto
	@# Generate go, gateway
	cd $(<D) && protoc \
		-I. \
		-I.. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mmanagement.proto=go.nlx.io/nlx/management-api/api,\
plugins=grpc:. \
		$(<F)

api/external/mock/mock_external.go: api/external/external.pb.go
	@# Generate mock file
	mockgen -source $< -destination $@

pkg/database/mock/mock_database.go: pkg/database/database.go
	@# Generate mock file
	mockgen -source $< -destination $@

pkg/directory/mock/mock_client.go: pkg/directory/client.go
	@# Generate mock file
	mockgen -destination $@ go.nlx.io/nlx/management-api/pkg/directory Client

pkg/management/mock/mock_client.go: pkg/management/client.go
	@# Generate mock file
	mockgen -destination $@ go.nlx.io/nlx/management-api/pkg/management Client

pkg/auditlog/mock/mock_auditlog.go: pkg/auditlog/logger.go
	@# Generate mock file
	mockgen -source $< -destination $@

pkg/txlogdb/mock/mock_database.go: pkg/txlogdb/database.go
	@# Generate mock file
	mockgen -source $< -destination $@

db/migrations.go:
	@# Generate migrations Go file
	go-bindata -prefix db/migrations -pkg db -o db/migrations.go db/migrations/...
