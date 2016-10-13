.PHONY: build

TEST_CONTAINER_PORT=8991
TEST_CONTAINER_NAME=auth-server
TEST_IMAGE_NAME=auth-server-image
EXPOSED_CONTAINER_PORT=8001

build:
	go get -t -d -v ./web-server
	go build -v -o WebServer ./web-server
	go get -t -d -v ./auth-server
	go build -v -o AuthServer ./auth-server

build-image:
	docker build -t $(TEST_IMAGE_NAME) auth-server

run-container:
	docker run -d -p $(TEST_CONTAINER_PORT):$(EXPOSED_CONTAINER_PORT) --name $(TEST_CONTAINER_NAME) $(TEST_IMAGE_NAME)

kill-containers:
	docker ps | grep $(TEST_IMAGE_NAME) | awk '{print $$1}' | xargs docker stop
	docker ps -a | grep $(TEST_IMAGE_NAME) |  awk '{print $$1}' | xargs docker rm

test:
	- go test -v ./web-server

integration-test: build build-image run-container test kill-containers