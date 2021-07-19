sr	=	44100
kr	=	4410
ksmps	=	10
nchnls	=	2

instr	1

;p4 grain pointer (in seconds)
;p5 pan (0...1)


ipanl	table	1-p5 ,4,1
ipanr	table	p5 ,4,1

andx	line	p4,p3,p4+p3
asig	tablei	andx*sr,1

k1	oscil	30000,1/p3,2	
asig	=	asig*k1
	
	outs	asig*ipanl, asig*ipanr
	
endin
