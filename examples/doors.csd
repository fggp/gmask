<CsoundSynthesizer>
; adapted from Andre Bartetzki's original cmask example
; see http://www.bartetzki.de/en/index.html
<CsOptions>
  -d -o dac
</CsOptions>

<CsInstruments>
sr     = 44100
ksmps  = 10
nchnls = 2
0dbfs  = 1.0

garev   init   0

  instr  1
;p4 transposition (1=normal)
;p5 table number (1...6)
;p6 pan (0...1)
;p7 dry/wet (0...1)

ipanl = table(1-p6, 10, 1)
ipanr = table(p6, 10, 1)

k1 = expon(0.5, p3, 0.01)
a1   loscil k1, p4, p5, 1, 0, 0, 2
a1 = linen:a(a1, 0, p3, 0.05)

garev = garev + a1*p7  
a2    = a1*ipanr
a1    = a1*ipanl
     outs  a1*(1-p7*p7), a2*(1-p7*p7)
  endin

  instr 99
krev = expseg(0.03, p3-5, 4, 5, 4) 
kral = linseg(0, p3*0.3, 1.1, p3*0.3, 0, p3*0.4, 0)
kral = kral*kral  

a1  = alpass(garev, kral, 0.05)
a2  = alpass(garev, kral, 0.06)
a1  = a1 * kral
a2  = a2 * kral
a1r = reverb2(garev+a1, krev, 0.3)
a2r = reverb2(garev+a2, krev*1.2, 0.33)
     outs  a1r+a1/2, a2r+a2/2

garev  =   0
  endin
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 0 -1 "door1.aif" 0 4 1
f2 0 0 -1 "door2.aif" 0 4 1
f3 0 0 -1 "door3.aif" 0 4 1
f4 0 0 -1 "door4.aif" 0 4 1
f5 0 0 -1 "door5.aif" 0 4 1
f6 0 0 -1 "door6.aif" 0 4 1

f10 0 8192 9 .25 1 0
i99 0 27
}

f 0 20

p1 const 1

p2 rnd beta .05 .1
mask (12 .01 18 .2) (12 .1 18 1)

p3 seg [.3 1.2 ipl 0.4]

p4 rnd uni mask [3 .8 ipl .4] [5 1.2 ipl .4]

p5 range 1 6
prec 0

p6 range 0 1

p7 seg (2 0 18 .5 ipl 1)
</CsScore>
</CsoundSynthesizer>
