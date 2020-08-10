	aseg
	org	fe06h

	ld	a, c
	cp	2
	jr	z, putchar
	cp	9
	jr	z, putstr
	halt

putchar:
	ld	a, e
	out	(0), a
	ret

putstr:
	ld	a, (de)
	cp	'$'
	ret	z
	out	(0), a
	inc	de
	jr	putstr
