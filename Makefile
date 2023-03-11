exec := "uhash"

.PHONY: build
# Build app for current operating system
build::
	@mkdir -p build
	@go build -o build/${exec} .
	@chmod +x build/${exec}

.PHONY: test
test::
	python test.py
	