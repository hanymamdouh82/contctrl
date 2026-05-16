BINARY := contctrl
INSTALL_DIR := /usr/local/bin/

.PHONY: build install

build:
	go build -o $(BINARY) ./cmd/main.go

install: build
	sudo mv $(BINARY) $(INSTALL_DIR)/$(BINARY)
