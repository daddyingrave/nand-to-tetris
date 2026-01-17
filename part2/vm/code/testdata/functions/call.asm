// call: callee from caller, nArgs 2, file: file

// return address label
  @file.caller.$ret.1
  D=A
  @SP
  A=M
  M=D
  D=A+1
  @SP
  M=D

// save callers segments
// push LCL
  @LCL
  D=M
  @SP
  A=M
  M=D
  D=A+1
  @SP
  M=D

// push ARG
  @ARG
  D=M
  @SP
  A=M
  M=D
  D=A+1
  @SP
  M=D

// push THIS
  @THIS
  D=M
  @SP
  A=M
  M=D
  D=A+1
  @SP
  M=D

// push THAT
  @THAT
  D=M
  @SP
  A=M
  M=D
  D=A+1
  @SP
  M=D

// Reposition ARG
  @SP
  D=M
  @5
  D=D-A
  @2
  D=D-A
  @ARG
  M=D

// Reposition LCL
  @SP
  D=M
  @LCL
  M=D

// goto function
  @callee
  0;JMP
(file.caller.$ret.1)

