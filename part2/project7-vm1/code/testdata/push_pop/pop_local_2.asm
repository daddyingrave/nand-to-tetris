// pop segment: local, index 2

@2
D=A
@LCL
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

