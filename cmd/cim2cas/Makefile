%.cas: %.cim
	go run cim2cas.go -cim $< -cas $@

%.cim: %.asm
	zmac --od . --oo cim,lst $<

.PHONY: clean
clean:
	rm -f *.cim *.lst
	rm -f *.cas
	rm -f cim2cas cim2cas.exe
