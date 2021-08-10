BIN="./bin"
SRC=$(shell find . -name "*.go")


.PHONY: go mocks

default: all

all: go mocks

go:
	$(info ******************** compiling proto to go ********************)
	protoc \
	--proto_path=./v1alpha1 \
	--go_out=./pkg/v1alpha1 \
	--go_opt=paths=source_relative \
	--go_opt=Mv1alpha1/kubelt.proto=github.com/proofzero/kubelt/proto/v1alpha1 \
	--go-grpc_out=./pkg/v1alpha1 \
	--go-grpc_opt=paths=source_relative \
	v1alpha1/kubelt.proto
	$(info )

mocks:
	$(info ******************** generating mocks ********************)
	GO111MODULE=on
	mockgen -source="pkg/v1alpha1/kubelt_grpc.pb.go" -package="v1alpha1" > "pkg/v1alpha1/test/kubelt_grpc_mocks.go"
	$(info )
