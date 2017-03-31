# zog
Z80 diassembler/assembler/emulator in golang

Just revisiting the first bit of serious programming I ever did - a Z80 disassembler.

This code currently (March 2017) implements most of the instruction decode logic (missing the ED prefix at the moment)
with the ability for instructions to represent themselves as strings, giving the basics of a disassembler (no labels
or directives).

Current WIP is adding an assembler (using the github.com/pointlander/peg parser), so that the emulator tests can be written in assembly rather than machine code.
