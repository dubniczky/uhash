exec := "uhash"

.PHONY: build
# Build app for current operating system
build::
	@mkdir -p build
	@go build -o build/${exec} .
	@chmod +x build/${exec}

.PHONY: test
# Run end-to-end tests using ptyhon
test::
	@python test.py

.PHONY: publish
# Create publish releases for all operating systems
publish::
	@mkdir -p build && rm -rf build/*
	GOOS=windows GOARCH=amd64 go build -o build/win64-${exec}.exe .
	GOOS=linux GOARCH=amd64 go build -o build/lin64-${exec} .
	GOOS=darwin GOARCH=amd64 go build -o build/mac64-${exec} .

	@chmod +x build/*
	@mkdir -p dist && rm -rf dist/*
	tar -czf dist/${exec}-linux-amd64.tar.gz -C build lin64-${exec}
	tar -czf dist/${exec}-darwin-amd64.tar.gz -C build mac64-${exec}
	tar -czf dist/${exec}-windows-amd64.tar.gz -C build win64-${exec}.exe
