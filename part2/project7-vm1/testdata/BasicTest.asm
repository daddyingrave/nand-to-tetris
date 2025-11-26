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

// push segment: constant, index 3040

@3040
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// pop segment: pointer, index 1

@THAT
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

// push segment: constant, index 32

@32
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// pop segment: this, index 2

@THIS
D=M
@2
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

// push segment: constant, index 46

@46
D=A
@SP
A=M
M=D
D=A+1
@SP
M=D

// pop segment: that, index 6

@THAT
D=M
@6
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

// push segment: pointer, index 0

@THIS
D=M
@SP
A=M
M=D
D=A+1
@SP
M=D

// push segment: pointer, index 1

@THAT
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

// push segment: this, index 2

@THIS
D=M
@2
A=D+A
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

// push segment: that, index 6

@THAT
D=M
@6
A=D+A
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

