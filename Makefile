.PHONY: build
build:
	go build -v -i

.PHONY: test
test:
	go test -v ./...

.PHONY: tags
tags:
	gotags -f tags -R .
