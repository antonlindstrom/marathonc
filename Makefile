all: get-deps build

build:
	@mkdir -p bin/
	@go build -o bin/marathonc marathonc.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

format:
	@go fmt ./...

clean:
	@rm -rf bin/
