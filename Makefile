.PHONY: build

build:
	bash -c "cd web-server; go get -t -d -v ./... && go build -v ./..."
	bash -c "cd auth-server; go get -t -d -v ./... && go build -v -o AuthServer ./..."
