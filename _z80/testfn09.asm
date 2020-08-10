	aseg
	org	100h

	ld	c, 9
	ld	de, msg
	call	5
	halt

msg:
	db	'Hello 09h$'
