BIN="./bin"
SRC=$(shell find . -name "*.go")


.PHONY: spec proto go mocks

default: all

all: spec proto go mocks

spec:
	$(info ******************** compiling open api spec ********************)
	swagger-cli bundle ./openapi/v1alpha1/rpc.yaml --outfile=openapi/v1alpha1/kubelt.yaml --type yaml
	$(info )

proto:
	$(info ******************** compiling spec to proto ********************)
	gnostic --resolve-refs --grpc-out=v1alpha1 openapi/v1alpha1/kubelt.yaml
	sed -i '4i\option go_package="kubelt.com/proto/v1alpha1";' v1alpha1/kubelt.proto
	$(info )

go:
	$(info ******************** compiling proto to go ********************)
	protoc \
	--proto_path=./v1alpha1 \
	--go_out=./pkg/v1alpha1 \
	--go_opt=paths=source_relative \
	--go_opt=Mv1alpha1/kubelt.proto=kubelt.com/proto/v1alpha1 \
	--go-grpc_out=./pkg/v1alpha1 \
	--go-grpc_opt=paths=source_relative \
	v1alpha1/kubelt.proto
	$(info )

mocks:
	$(info ******************** generating mocks ********************)
	GO111MODULE=on
	mockgen -source="pkg/v1alpha1/kubelt_grpc.pb.go" -package="v1alpha1" > "pkg/v1alpha1/test/kubelt_grpc_mocks.go"
	$(info )
