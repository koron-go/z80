.PHONY: run
run: zexdoc.cim
	go run .

.PHONY: profile
profile: zexdoc.cim
	go run . -cpuprof cpu.pprof
	go tool pprof -svg ./cpu.pprof > cpuprof.svg

%.cim: %.asm
	go generate

.PHONY: clean
clean:
	rm -f zexdoc.lst
	rm -f cpu.pprof
	rm -f cpuprof.svg
