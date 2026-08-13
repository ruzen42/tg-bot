BUILD_DIR=bin
VERSION=0.1.1
NAME=bot
TOKEN != cat .env

.PHONY: all build clean run fmt

all: build

build: fmt
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BUILD_DIR)/$(NAME) .

clean:
	rm -rf $(BUILD_DIR)

run: build
	export $(TOKEN) 
	./$(BUILD_DIR)/$(NAME)

fmt:
	go fmt
