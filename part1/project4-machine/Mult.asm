// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

    @R0
    D=M
    @multiplicand
    M=D

    @R1
    D=M
    // We will just sum up multiplicand multiplicator times
    @multiplicator
    M=D

    @R2
    M=0

(LOOP)
    // Decrement multiplicator
    @multiplicator
    D=M-1
    M=D
    // If multiplicator less than 0 we done
    @END
    D;JLT
    
    // Add multiplicand to R2 accumulator
    @multiplicand
    D=M
    @R2
    D=D+M
    M=D

    @LOOP
    0;JMP

(END)
    @END
    0;JMP
