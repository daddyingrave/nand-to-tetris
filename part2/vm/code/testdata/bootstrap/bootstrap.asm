// bootstrap start

  @256
  D=A
  @SP
  M=D
// call: Sys.init from bootstrap, nArgs 0, file: file

// return address label
  @file.bootstrap.$ret.1
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
  @0
  D=D-A
  @ARG
  M=D

// Reposition LCL
  @SP
  D=M
  @LCL
  M=D

// goto function
  @Sys.init
  0;JMP
(file.bootstrap.$ret.1)



// bootstrap ends

