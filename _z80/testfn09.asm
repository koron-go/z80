	aseg
	org	100h

	ld	c, 9
	ld	de, msg
	call	5
	jp	0

msg:
	db	'Hello 09h$'
