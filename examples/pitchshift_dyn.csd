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
0dbfs  = 1.0

  instr 1

;p2 onset
;p3 duration
;p4 speed factor (=transposition)

kenv  = oscil(0.92, 1/p3, 2)
aindx = line:a(p2, p3, p2 + p3*p4)
asig  = tablei:a(aindx*sr, 1)
     out asig*kenv
  endin  
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 131072 1 "schwermt.aif" 0 4 1  
f2 0 8193 8 0 4096 1 4096 0
}

; for use with pitchshift.orc !

f 0 2.2  

p1 const 1

p2 const 0.01   ;constant grain interonset 10 ms

p3 const 0.02   ;constant grain duration 20 ms

p4 seg [1 2.2 ipl .3]
      ;acceleration = glissando
</CsScore>
</CsoundSynthesizer>
