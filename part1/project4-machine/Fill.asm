// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

   @SCREEN
   D=A
   @addr
   M=D

   @start
   M=D

   @KBD
   D=A
   @end
   M=D

   @brush
   M=0

(LOOP)
   @KBD
   D=M
   @BLACK
   D;JGT
   @WHITE
   D;JEQ

// Choosing the black brush
(BLACK)
   @brush
   M=-1
   @FILL
   0;JMP

// Choosing the white brush
(WHITE)
   @brush
   M=0
   @FILL
   0;JMP

// Brushing
(FILL)
   @brush
   D=M

   // Assign next addres to @addr
   @addr
   A=M
   M=D
   D=A+1
   @addr
   M=D

   // continue
   @end
   D=M-D
   @LOOP
   D;JNE
   
   // reset 
   @start
   D=M
   @addr
   M=D
   @LOOP
   M;JMP

(END)
   @END
   0;JMP



