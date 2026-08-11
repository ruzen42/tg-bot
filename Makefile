GC=go
BUILD_DIR=output
VERSION=0.1.0
NAME=bot

.PHONY: all build clean run fmt

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BUILD_DIR)/$(NAME) .

clean:
	rm -rf $(BUILD_DIR)

run: build
	export $(cat .env)
	./$(BUILD_DIR)/$(NAME)

fmt:
	go fmt