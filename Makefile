BIN=peerwatch
TARGET=$(CURDIR)/bin/$(BIN)
SRC=$(CURDIR)/src/*
STATUS= "\033[5m\033[32mDONE\033[0m \n"

.PHONY: all

all: build install clean

build:
	@echo -n "\n[\033[5m\033[32m+\033[0m] Creating a build..."
	@go build -mod vendor -ldflags="-s -w" -o $(TARGET) $(SRC)
	@echo $(STATUS)

clean:	
	@echo -n "\n[\033[5m\033[32m+\033[0m] Cleaning..."
	@rm -rf $(TARGET)
	@go clean -testcache -modcache -r -i 
	@echo $(STATUS)

install:
	@echo -n "\n[\033[5m\033[32m+\033[0m] Installing..."
	@cp $(TARGET) /usr/local/bin
	@echo $(STATUS)

uninstall: 
	@echo -n "\n[\033[5m\033[32m+\033[0m] Uninstalling..."
	@rm -f /usr/local/bin/$(BIN)
	@echo $(STATUS)
