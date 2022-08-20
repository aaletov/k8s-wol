include variables.mk

build-server:
	CGO_ENABLED=0 go build -o ./build/server/server ./cmd/server

docker-build-server:
	cp ./server/Dockerfile ./build/server/Dockerfile
	docker build --file ./build/server/Dockerfile --tag server:latest ./build/server 

kind-install:
	chmod +x $(KIND_DIR)/kind-install.sh
	$(KIND_DIR)/kind-install.sh

kind-create-cluster:
	${KIND} create cluster --config ./test/deploy/kind/kind.yaml --image kindest/node:v1.23.5 --wait 60s

kind-load-images:
	${KIND} load docker-image server:latest

run:
	go run ./main A8:A1:59:2F:6E:54

dependencies:
	go get k8s.io/apimachinery/pkg/apis/meta/v1
	go get k8s.io/client-go/kubernetes
	go get k8s.io/client-go/rest
	go get github.com/sirupsen/logrus