.PHONY: run
run: zexdoc.cim
	go run .

run-all: zexall.cim
	go run . -all

.PHONY: profile
profile: zexdoc.cim
	go run . -cpuprof cpu.pprof -memprof mem.pprof
	go tool pprof -svg ./cpu.pprof > cpuprof.svg

%.cim: ../../_z80/%.asm
	go generate

.PHONY: clean
clean:
	rm -f zexdoc.lst zexall.lst
	rm -f *.pprof
	rm -f cpuprof.svg
