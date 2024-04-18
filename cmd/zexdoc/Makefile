TIMESTAMP = $(shell date +%Y%m%dT%H%M%S%z)
PROFILE_NAME := tmp/profile/zexdoc-$(TIMESTAMP)

.PHONY: run
run: zexdoc.cim
	go run .

run-all: zexall.cim
	go run . -all

.PHONY: profile
profile: zexdoc.cim
	mkdir -p tmp/profile
	date -Ins > $(PROFILE_NAME).log
	go run . -cpuprof cpu.pprof -memprof mem.pprof | tee -a $(PROFILE_NAME).log
	date -Ins >> $(PROFILE_NAME).log
	go tool pprof -svg ./cpu.pprof > $(PROFILE_NAME).svg
	cp $(PROFILE_NAME).log profile.log
	cp $(PROFILE_NAME).svg cpuprof.svg

%.cim: ../../_z80/%.asm
	go generate

.PHONY: clean
clean:
	rm -f zexdoc.lst zexall.lst
	rm -f *.pprof
	rm -f cpuprof.svg
