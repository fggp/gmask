sr	=	44100
kr	=	4410
ksmps	=	10
nchnls	=	2

garev 	init 	0

instr	1

;p4 transposition (1=normal)
;p5 table number (1...6)
;p6 pan (0...1)
;p7 dry/wet (0...1)

ipanl	table	1-p6 ,10,1
ipanr	table	p6 ,10,1

k1	expon	.5,p3,.01
a1	loscil	k1,p4,p5,1,0,0,2
a1	linen	a1,0,p3,.05

garev	= garev + a1*p7	
a2	= 	a1*ipanr
a1	=	a1*ipanl
	outs	a1*(1-p7*p7),a2*(1-p7*p7)

endin

instr 99

krev	expseg	.03,p3-5,4,5,4	
kral	linseg	0,p3*.3,1.1,p3*.3,0,p3*.4,0
kral	= kral*kral	

a1	alpass	garev, kral,.05
a2	alpass	garev, kral,.06
a1	= a1 * kral
a2	= a2 * kral
a1r	reverb2	garev+a1,krev,.3
a2r	reverb2	garev+a2,krev*1.2,.33
	outs	a1r+a1/2,a2r+a2/2

garev	= 	0

endin	
