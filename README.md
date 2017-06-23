# zog

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
