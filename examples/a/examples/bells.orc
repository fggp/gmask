sr = 44100
kr = 4410
nchnls = 2

instr 1

;p2 onset
;p3 duration
;p4 base frequency
;p5 fm index
;p6 pan (L=0, R=1)

kenv	expon		1,p3,0.01
kindx	expon		p5,p3,.4
a1	foscil	kenv*10000,p4,1,1.143,kindx,1
	outs		a1*(1-p6),a1*p6
		
endin	
