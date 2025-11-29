// gt

  @SP
  A=M
  A=A-1
  D=M
  A=A-1
  D=M-D
  @GT_666
  D;JGT
  @NOT_GT_666
  D;JMP
(GT_666)
  D=-1
  @GT_END_666
  D;JMP
(NOT_GT_666)
  D=0
(GT_END_666)
  @SP
  A=M
  A=A-1
  A=A-1
  M=D
  D=A+1
  @SP
  M=D

