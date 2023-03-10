exec := "uhash"

.PHONY: build
# Build app for current operating system
build::
	@mkdir -p build
	@go build -o build/${exec} .
	@chmod +x build/${exec}

.PHONY: test
test::
	@make build
	@echo "--- sha1 a"
	@./build/${exec} sha1 a
	@echo "--- md5 a"
	@./build/${exec} md5 a
	@echo "--- sha-256 a"
	@./build/${exec} sha-256 a