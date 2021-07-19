sr	=	44100
kr	=	4410
ksmps	=	10
nchnls	=	2

instr	1

;p4 frequency
;p5 pan (0...1) 

ipanl	table	1-p5 ,1,1
ipanr	table	p5 ,1,1

k1	expon	1,p3,.01
a1	foscil	k1*4200,p4,1,2.41,k1*6,2
	
	outs	a1*ipanl, a1*ipanr
	
endin
