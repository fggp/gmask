sr = 44100
kr = 4410
nchnls = 2

instr 1		;mallet ?

;p2 onset
;p3 duration
;p4 pitch (0-4)
;p5 octav (7-9)

kenv	oscil		1,1/p3,2
kindx	kpow		kenv,6,.5
iton	table		p4,5
a1	foscil	kenv*8000,cpspch(p5+iton),1,4,kindx,1
	outs		a1*(1-p4/4),a1*p4/4
		
endin	

instr 2		;metal plate

;p2 onset
;p3 duration
;p4 pitch (0/1)

kindx	expon		1,p3,.001
a1 	rand		100
a2	oscil		10000*kindx,3000+1500*p4+a1*(1+kindx),1

	outs		a2*p4,a2*(1-p4)	
endin	

instr 3

;p2 onset
;p3 duration
;p4 pitch (0-3)

kenv	oscil		1,1/p3,3
kindx	oscil		2,1/p3,4
a1	foscil	kenv*11000,100+p4*20,1,1.4,kindx,1
	outs		a1,a1
	
endin	
