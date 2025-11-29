// push segment: constant, index 111

@111
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// push segment: constant, index 333

@333
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// push segment: constant, index 888

@888
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// pop segment: static, index 8

@BasicTest.8
D=A
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

// pop segment: static, index 3

@BasicTest.3
D=A
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

// pop segment: static, index 1

@BasicTest.1
D=A
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

// push segment: static, index 3

@BasicTest.3
D=M
@SP
A=M
M=D
D=A+1
@SP
M=D

// push segment: static, index 1

@BasicTest.1
D=M
@SP
A=M
M=D
D=A+1
@SP
M=D

// sub

@SP
A=M
A=A-1
D=M
A=A-1
D=M-D
@SP
A=M
A=A-1
A=A-1
M=D
D=A+1
@SP
M=D

// push segment: static, index 8

@BasicTest.8
D=M
@SP
A=M
M=D
D=A+1
@SP
M=D

// add

@SP
A=M
A=A-1
D=M
A=A-1
D=D+M
@SP
A=M
A=A-1
A=A-1
M=D
D=A+1
@SP
M=D

