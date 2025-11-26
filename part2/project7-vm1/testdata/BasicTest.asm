// push segment: constant, index 3030

@3030
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// pop segment: pointer, index 0

@THIS
D=M
@0
D=D+A
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

