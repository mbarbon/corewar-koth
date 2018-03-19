;redcode
;name Commando
;author A. K. Dewdney
;strategy ICWST'86 finalist
;strategy the original!
;assert 1
        MOV     #0     ,-1
        JMP     -1
START   SPL     -1              ; changed from SPL -2 to correct timing -
                                ; SPL behaves differently under '86 & '88
        MOV     IMP    ,123
        SPL     122
COPY    MOV     <LOC   ,<NEW
        CMP     #-12   ,LOC     ; was CMP LOC ,#-12  ('88 violation)
        JMP     COPY
        JMP     106
IMP     MOV     0      ,1
NEW     DAT             #113
LOC     DAT             #0
        END     START
