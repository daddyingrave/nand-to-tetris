// pop segment: argument, index 1

@1
D=A
@ARG
D=D+M
@R13
M=D
@SP
A=M-1
D=M
@R13
A=M
M=D
@SP
M=M-1

