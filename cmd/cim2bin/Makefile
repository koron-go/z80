%.bin: %.cim
	go run cim2bin.go -cim $< -bin $@

%.cim: %.asm
	zmac --od . --oo cim,lst $<

.PHONY: clean
clean:
	rm -f *.cim *.lst
	rm -f *.bin
	rm -f cim2bin cim2bin.exe
