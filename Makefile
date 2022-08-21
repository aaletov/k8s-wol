include variables.mk

build-server:
	CGO_ENABLED=0 go build -o ./build/server/server ./cmd/server

docker-build-server:
	cp ./server/Dockerfile ./build/server/Dockerfile
	docker build --file ./build/server/Dockerfile --tag server:latest ./build/server 

kind-install:
	go install sigs.k8s.io/kind@${KIND_VERSION}

install-protoc:
	sudo apt install -y protobuf-compiler

compile-proto:
	mkdir -p api/generated/v1/
	protoc -I=api/v1 --go-grpc_out=api/generated api/v1/*.proto
	protoc -I=api/v1 --go_out=api/generated api/v1/*.proto

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate-deepcopy:
	controller-gen object paths=./api/generated/v1 output:dir=./api/generated/v1

kind-create-cluster:
	${KIND} create cluster --config ./test/deploy/kind/kind.yaml --image kindest/node:${KIND_NODE_TAG} --wait 60s

kind-load-images:
	${KIND} load docker-image server:latest

run:
	go run ./main A8:A1:59:2F:6E:54

dependencies:
	go get k8s.io/apimachinery/pkg/apis/meta/v1
	go get k8s.io/client-go/kubernetes
	go get k8s.io/client-go/rest
	go get github.com/sirupsen/logrus