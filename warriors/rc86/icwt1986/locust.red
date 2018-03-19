;redcode
;name Locust
;author Mark Clarkson
;strategy ICWST'86 finalist
;assert 1
START   MOV     IMP    ,NEW
        MOV     #0     ,-2
        SPL     NEW
        JMP     START
IMP     MOV     0      ,1
A       DAT             #0
B       DAT             #0
E       DAT             #0
NEW     DAT             #0
        END     START
