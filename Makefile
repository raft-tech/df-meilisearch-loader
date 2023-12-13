n.PHONY: build, delete, deploy, docker, install, kind, fmt, vet, test

# Remember to do "export DF_HOME=/path/to/data-fabric/" before running make build
PROJECT_HOME=${PWD}
REGISTRY=ghcr.io/raft-tech
IMAGE=df-meili
VERSION=dev
FULL_IMAGE=${REGISTRY}/${IMAGE}:${VERSION}
KIND_CLUSTER=data-fabric

build: fmt vet
	go build -o ${IMAGE} ./cmd

deploy:
	kubectl apply -f ${PROJECT_HOME}/deploy.yaml

docker:
	docker build --no-cache -f ${PROJECT_HOME}/Dockerfile \
       ${PROJECT_HOME}/ \
       -t ${FULL_IMAGE}

ghcr:
	docker buildx build \
	   --platform linux/amd64 \
	   -f ${PROJECT_HOME}/Dockerfile \
       ${PROJECT_HOME}/ \
       -t ${FULL_IMAGE} \
	   --push 

kind:
	kind load docker-image ${FULL_IMAGE} --name ${KIND_CLUSTER}

pull:
	docker pull ${FULL_IMAGE}

push:
	docker push ${FULL_IMAGE}

fmt:
	go fmt ./...

vet:
	go vet ./...

test: vet
	go clean -cache
	go test ./... -v
