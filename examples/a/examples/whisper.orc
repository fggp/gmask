sr = 44100
kr = 4410
nchnls = 2

instr 1

ipanl	table	1-p5 ,4,1
ipanr	table	p5 ,4,1

andx	line	p4,p3,p4+p3*p6
asig	tablei	andx*sr,1
kamp	oscil	8000,1/p3,2
		outs	asig*kamp*ipanl, asig*kamp*ipanr  
	
endin	
