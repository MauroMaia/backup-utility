COMMIT_ID := $(shell git rev-list --tags --max-count=1)
VERSION := $(shell git describe --abbrev=0 --tags --match "v[0-9]*" $(COMMIT_ID))
BUILD_DATE := $(shell date --iso=s)

COLOR_RED := $(shell echo -e "\033[0;31m")
COLOR_YELLOW := $(shell echo -e "\033[0;33m")
COLOR_END := $(shell echo -e "\033[0m")

all: clean _version test build
	@echo "all done"

build: clean _version
	@echo -e "$(COLOR_YELLOW)Setting up build folders $(COLOR_END)"
	mkdir build || true
	@echo -e "$(COLOR_YELLOW)Setting in app version $(COLOR_END)"
	sed -i "s/\(.*VERSION = \).*/\1 \"$(VERSION)\"/g" src/main.go
	sed -i "s/\(.*COMMIT_ID = \).*/\1 \"$(COMMIT_ID)\"/g" src/main.go
	sed -i "s/\(.*BUILD_DATE = \).*/\1 \"$(BUILD_DATE)\"/g" src/main.go
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	cd src && GOOS=linux GOARCH=amd64 go build -o ../build/backup-utility -v -race main.go

test:
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	@#go test --json -v  ./..
	 go test -v  ./src/...

_version:
	@echo -e "backup-utility version: $(VERSION)"

help: _version
	@echo -e "$(COLOR_YELLOW)List of actions that can be executed $(COLOR_END)"
	@echo " -> build"
	@echo " -> verison"
	@echo " -> help (this action)"

clean:
	@echo -e "$(COLOR_YELLOW)Removed old files $(COLOR_END)"
	rm build/backup-utility || true
