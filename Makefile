# Name of the binary to be created
BINARY_NAME=GuessTheNumber

# Path to the directory where the binary will be installed
INSTALL_PATH=${DESTDIR}/usr/bin

# Flags to pass to the go build command
GO_BUILD_FLAGS=-v

.DEFAULT_GOAL := build

.PHONY: deps
deps:
	go get -v

.PHONY: build
build: deps
	go build ${GO_BUILD_FLAGS}
	strip ${BINARY_NAME}

.PHONY: install
install:
	mkdir -p ${INSTALL_PATH}
	cp ${BINARY_NAME} ${INSTALL_PATH}

.PHONY: uninstall
uninstall:
	rm ${INSTALL_PATH}/${BINARY_NAME}
