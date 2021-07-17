Adaptation of Cmask to the Go programming language
========

This package is an adaptation of Andre Bartetzki's Cmask software for the Go
programming language. It provides a standalone program called gmask which uses the library gmasklib (https://github.com/fggp/gmasklib).

The lexer (gmlex.go) uses the Go text/scanner package which is enough for the
tokens used in the Cmask grammar
(see http://www2.ak.tu-berlin.de/~abartetzki/CMaskMan/CMask-Reference.htm).

The parser is defined in the gmask.y file. It is generated using the command
`go tool yacc -o gmask.go gmask.y`.

Then one can build and install the program gmask.

Once the program is built, you can run the examples in the 'examples' directory with the command 'csound example-name.csd', provided that the gmask program is in your path, and the samples are reachable, e.g. in SSDIR.

The original Cmask software was written by Andre Bartetzki. Andre has stopped Cmask development but the software is still present on his site: http://www.bartetzki.de/en/software.html.

Cmask was published under GPL. Thanks to Andre who kindly allowed me to publish gmask under LGPL:  

"Dear François,

thanks for bringing Cmask to a new life!  
Yes, you may publish Gmask under LGPL.  
best

Andre"

The gmask program reflects exactly Cmask features. It has a parser that recognizes the grammar written by Andre: http://www2.ak.tu-berlin.de/~abartetzki/CMaskMan/CMask-Reference.htm

When the program is called on a parameter file respecting Cmask language, it will output a
Csound sco file on standard out. One can also write the attribute bin="gmask" in a CsScore tag of a csd file to get the score generated on the fly while playing the csd file with Csound. See the examples directory in the gmask/gmask directory.
