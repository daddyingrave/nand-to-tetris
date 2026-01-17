// return for function: fn, file: file

// endFrame
  @LCL
  D=M
  @endFrame
  M=D

// retAddr
  @5
  D=D-A
  A=D
  D=M
  @retAddr
  M=D

// reposition return val
  @SP
  AM=M-1
  D=M
  @ARG
  A=M
  M=D

// reposition SP for caller
  @ARG
  D=M+1
  @SP
  M=D

// restore caller segments
  @1
  D=A
  @endFrame
  D=M-D
  A=D
  D=M
  @THAT
  M=D

  @2
  D=A
  @endFrame
  D=M-D
  A=D
  D=M
  @THIS
  M=D

  @3
  D=A
  @endFrame
  D=M-D
  A=D
  D=M
  @ARG
  M=D

  @4
  D=A
  @endFrame
  D=M-D
  A=D
  D=M
  @LCL
  M=D

// goto retAddr
  @retAddr
  A=M
  0;JMP

