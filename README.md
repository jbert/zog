# zog

Z80 diassembler/assembler/emulator in golang

Update: Sat 28 Oct 11:38:39 BST 2017

So the cpu core passes zexdoc and has a quick and dirty pass for timing
correct T-states. Similarly quick and dirty Spectrum screen and keyboard
support and 'z80' file format support are sufficient to:

1. boot the original spectrum 48K ROM to basic and run the Most Important
Program:

    10 PRINT "HELLO"

    20 GOTO 10


2. Load and run Manic Miner and Jetpac at what seem to be the correct speed.

3. Fail to run Elite, perhaps due to lacking interrupt mode support.


I think I'm unlikely to take it any further, since the world doesn't need
another spectrum emulator. However, this was a *lot* of fun and both rewarding
and challenging. I'd obviously do things quite a lot differently if I did it
again, but learning is kind of the point.

If by any chance you are here looking for a spectrum emulator, you will find
better ones elsewhere. If you are looking for how to code an emulator, you
will almost certainly find better code elsewhere. But if you are writing your
own and want to share thoughts, please let me know :-)


----------------


Z80 diassembler/assembler/emulator in golang

Just revisiting the first bit of serious programming I ever did - a Z80
disassembler.

This code currently (June 2017) implements all the instruction decode logic
and has a String() representation for instructions. There is also a basic
assembler (using the github.com/pointlander/peg parser), so that the emulator
tests can be written in assembly rather than machine code. This gives us
bytes -> string -> bytes conversions.

I had ambitions of wiring an assembler which could parse the
[zexall](http://mdfs.net/Software/Z80/Exerciser/) code as
a test suite, but a macro assembler is a bigger undertaking than I'd like to
get into prior to actually having an emulator.

Instead, the zmac assembler can be used to build zexall and next steps are
actually to write the emulator.
