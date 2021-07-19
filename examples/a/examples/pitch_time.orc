sr = 44100
kr = 4410

instr 1

;p2 onset
;p3 duration
;p4 sound file pointer
;p5 speed factor (=transposition)

kenv	oscil		20000,1/p3,4
aindx	line		p4,p3,p4+p3*p5
asig	tablei	aindx*sr,1

	out		asig*kenv
	
endin	
