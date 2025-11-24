// gt

@SP
A=M
A=A-1
D=M
A=A-1
D=M-D
@GT
D;JGT
@NOT_GT
D;JMP
(GT)
D=1
D=-1
@SP
A=M
A=A-1
M=D
D=A
@SP
@M=D

