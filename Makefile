GO_TOOLS=github.com/mitchellh/gox

# Determine the arch/os we're building for
GOX_ARCH=amd64
GOX_OS=darwin linux windows

BUILD_DIR=build


# all builds binaries for all targets
all: bin

bin: tools
	@echo "==> Removing old directory..."
	@rm -rf $(BUILD_DIR)/
	@mkdir -p $(BUILD_DIR)/

	@echo "==> Building..."
	@gox \
		-os="$(GOX_OS)" \
		-arch="$(GOX_ARCH)" \
		-osarch="!darwin/arm !darwin/arm64" \
		-output "$(BUILD_DIR)/{{.OS}}_{{.Arch}}/hr" \
		.

clean:
	@rm -rf $(BUILD_DIR)
	@go clean

uninstall:
	@go clean -i

tools:
	@go get -u $(GO_TOOLS)


.PHONY: all bin clean tools
