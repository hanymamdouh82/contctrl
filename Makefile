BINARY := contctrl
INSTALL_DIR := /home/hany/dotfiles/.bin/dev

.PHONY: build install

build:
	go build -o $(BINARY) ./cmd/main.go

install: build
	sudo mv $(BINARY) $(INSTALL_DIR)/$(BINARY)
