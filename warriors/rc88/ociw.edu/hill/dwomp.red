;redcode
;name DWOMP
;author MARK A. DURHAM
;assert 1
       SPL     DWARF
       JMP     STOMP
INC    DAT             #2396
STOMP  JMP     STOMP,  <-2
BOMB   DAT             #0
DWARF  ADD     INC,    BOMB
       MOV     BOMB,   @BOMB
       JMP     DWARF
       END

