<CsoundSynthesizer>
; adapted from Andre Bartetzki's original cmask example
; see http://www.bartetzki.de/en/index.html
<CsOptions>
  -d -o dac
</CsOptions>

<CsInstruments>
sr     = 44100
ksmps  = 10
nchnls = 1

instr 1

;p2 onset
;p3 duration
;p4 sound file pointer

kenv	oscil		20000,1/p3,2
aindx	line		p4,p3,p3+p4
asig	tablei	aindx*sr,1

	out		asig*kenv

endin	
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 131072 1 "schwermt.aif" 0 4 1  ;43520
f2 0 8193 8 0 4096 1 4096 0
}

; for use with timestretch.orc !

f 0 11  

p1 const 1

p2 range .01 .05	;random grain interonset 10 ... 50 ms
prec 2

p3 range .04 .1 	;random grain duration 40 ... 100 ms
prec 2

p4 range 0 2.1	;random pointer
prec 3
</CsScore>
</CsoundSynthesizer>

