// eq

@SP
A=M
A=A-1
D=M
A=A-1
D=D&M
D=!D
D=D&M
@EQ
D;JEQ
@NOT_EQ
D;JNE
(EQ)
D=1
(NOT_EQ)
D=-1
@SP
A=M
A=A-1
M=D
D=A
@SP
M=D

