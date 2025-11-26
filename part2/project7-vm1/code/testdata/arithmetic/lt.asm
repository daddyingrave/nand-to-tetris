// lt

@SP
A=M
A=A-1
D=M
A=A-1
D=M-D
@LT
D;JLT
@NOT_LT
D;JMP
(LT)
D=1
(NOT_LT)
D=-1
@SP
A=M
A=A-1
A=A-1
M=D
D=A+1
@SP
M=D

