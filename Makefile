.PHONY: build
build:
	go build -gcflags '-e'

.PHONY: test
test:
	go test ./...

.PHONY: tags
tags:
	gotags -f tags -R .

.PHONY: cover
cover:
	mkdir -p tmp
	go test -coverprofile tmp/_cover.out ./...
	go tool cover -html tmp/_cover.out -o tmp/cover.html

.PHONY: checkall
checkall: vet lint staticcheck

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: clean
clean:
	go clean
	rm -f tags
	rm -f tmp/_cover.out tmp/cover.html
	rm -f z80.asm

.PHONY: zexdoc
zexdoc:
	$(MAKE) -C cmd/zexdoc run

z80.asm:
	go tool compile -o z80.o -S `go list -f '{{join .GoFiles " "}}'` > $@
	rm -f z80.o

# based on: github.com/koron-go/_skeleton/Makefile
