.PHONY: run
run: zexdoc.cim
	go run .

.PHONY: profile
profile: zexdoc.cim
	go run . -cpuprof cpu.pprof -memprof mem.pprof
	go tool pprof -svg ./cpu.pprof > cpuprof.svg

%.cim: %.asm
	go generate

.PHONY: clean
clean:
	rm -f zexdoc.lst
	rm -f *.pprof
	rm -f cpuprof.svg
