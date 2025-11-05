.default: build

.PHONY: build
build:
	@go build -o bin/envet .

.PHONY: install
install:
	@go install .
