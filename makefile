all: format build

format:
	@go fmt .

build:
	@go build .