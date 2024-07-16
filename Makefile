CLIENT_DIR := client
SERVER_DIR := server
CLIENT_BINARY := client-app
SERVER_BINARY := server-app


all: build-client build-server

build-client:
	go build -o $(CLIENT_DIR)/$(CLIENT_BINARY) $(CLIENT_DIR)/main.go

build-server:
	go build -o $(SERVER_DIR)/$(SERVER_BINARY) $(SERVER_DIR)/main.go

run-client: build-client
	./$(CLIENT_DIR)/$(CLIENT_BINARY)

run-server: build-server
	./$(SERVER_DIR)/$(SERVER_BINARY)

clean:
	rm -f $(CLIENT_DIR)/$(CLIENT_BINARY) $(SERVER_DIR)/$(SERVER_BINARY)


.PHONY: all build-client build-server run-client run-server clean