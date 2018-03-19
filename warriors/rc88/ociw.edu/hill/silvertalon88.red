;redcode 
;name Silver Talon '88
;author Ben Ford 
;strategy Blur '88 + Silver Talon 
;strategy with a twist to the stun 
;assert CORESIZE==8000 

step EQU 70 

top: mov ptr,<ptr 
a: add inc,scan 
scan: cmp -3*step+5,-3*step 
mov scan,@-3 
jmn top,@-3 
bomb: spl 0,<1-step 
mov inc,<bomb-2 
djn -1,@-1-step 
inc: dat <-step,<-step 
ptr: spl -1,#-2*step 

END scan 

