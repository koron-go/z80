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

.PHONY: zexdoc
zexdoc:
	$(MAKE) -C cmd/zexdoc run

switch.go: op_*.go gen_switch.go ./cmd/gen_switch/*.go
	rm -f switch.go switch.go.new
	go run ./cmd/gen_switch | goimports > switch.go.new
	mv switch.go.new switch.go

# based on: github.com/koron-go/_skeleton/Makefile
