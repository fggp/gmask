sr = 44100
kr = 4410
nchnls = 2

instr 1

k1	oscil	8000*p5,1/p3,1
a1	oscil	k1,p4,2
	outs	a1*(1-p6),a1*p6
	
endin	
