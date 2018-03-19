;redcode
;name Facta Non Verba
;author Michael Giberson
;strategy ICWST'86 finalist
;assert 1
TOP     MOV     IMP    ,900
LOOP    MOV     #0     ,TOP
        MOV     #0     ,TOP
        MOV     #0     ,TOP
        MOV     #0     ,TOP
        SPL     895
        JMP     LOOP
IMP     MOV     0      ,1
        END     TOP
