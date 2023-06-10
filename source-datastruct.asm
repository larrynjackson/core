;some stuff that's not complete yet. work in progress

ZERO:           DBYTE   0         
ONE:            DBYTE   1
TWO:            DBYTE   2
THREE:          DBYTE   3
FOUR:           DBYTE   4
FIVE:           DBYTE   5
SIX:            DBYTE   6
SEVEN:          DBYTE   7
EIGHT:          DBYTE   8

R0IDX:          DBYTE   2
R1IDX:          DBYTE   3
R2IDX:          DBYTE   4
R3IDX:          DBYTE   5
R4IDX:          DBYTE   6
R5IDX:          DBYTE   7
R6IDX:          DBYTE   8
R7IDX:          DBYTE   9

stk-a:pt:       dbyte   239
stk-a:lo:       dbyte   240
stk-a:hi:       dbyte   216

stkbuf:lo:      dbyte   255
stkbuf:hi:      dbyte   77



buffer:lo:      dbyte  255
buffer:hi:      dbyte  66


start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0

                    ldi     r7 stk-a-init:hi
                    shl     r7 r7 8
                    addi    r7 stk-a-init:lo
                    call    r7


                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo       
                    in      r0
            

                    ldi     r6 zero
                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo

pushBufNext:        ldw     r0 r7 r6
                    ldi     r5 pushDone:hi
                    shl     r5 r5 8
                    addi    r5 pushDone:lo
                    cmpi    r0 zero
                    bflag   r5 eq
                    
                    ldi     r3 stk-a-push:hi
                    shl     r3 r3 8
                    addi    r3 stk-a-push:lo
                    call    r3

                    addi    r6 one
                    ldi     r5 pushBufNext:hi
                    shl     r5 r5 8
                    addi    r5 pushBufNext:lo
                    jump    r5

pushDone:           ldi     r6 zero
                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo


popStackNext:       ldi     r5 stk-a-isempty:hi   
                    shl     r5 r5 8
                    addi    r5 stk-a-isempty:lo
                    call    r5
                    cmpi    r0 one

                    ldi     r5 endMain:hi
                    shl     r5 r5 8
                    addi    r5 endMain:lo
                    bflag   r5 eq

                    ldi     r5 stk-a-pop:hi
                    shl     r5 r5 8
                    addi    r5 stk-a-pop:lo
                    call    r5

                    stw     r0 r7 r6
                    addi    r6 one
                    ldi     r3 ZERO
                    stw     r3 r7 r6

                    ldi     r5 popStackNext:hi
                    shl     r5 r5 8
                    addi    r5 popStackNext:lo
                    jump    r5           
                 
                   


endMain:            ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo
                    out     r0

 endProgram:        halt




stk-a-init:         ldi     r0 stk-a:hi
                    shl     r0 r0 8
                    addi    r0 stk-a:pt
                    ldi     r1 zero
                    stw     r1 r0 r1
                    rtrn


                    ;pushes the contents of R0 onto the stack.
stk-a-push:         MVSR    r7                
                    ldi     r6 R0IDX
                    ldw     r0 r7 r6            ;R0 = value to push
          
                    ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:pt
                    ldi     r6 zero
                    ldw     r1 r7 r6            ;get stack-a top to R1
                    addi    r1 one
                    stw     r1 r7 r6            ;increment stack-a top

                    ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:lo
                    stw     r0 r7 r1            ;save r0 to stack-a[R1]
                    rtrn


                    ;returns the top of the stack value in R0
stk-a-pop:          ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:pt
                    ldi     r6 zero
                    ldw     r1 r7 r6             ;get stack-a top to R1

                    ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:lo
                    ldw     r0 r7 r1             ;R0 = stack-a[R1]

                    ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:pt
                    ldi     r6 zero
                    subi    r1 one               ;decrement stack-a top
                    stw     r1 r7 r6             ;save stack-a top

                    MVSR    r7
                    ldi     r6 r0idx
                    stw     r0 r7 r6             ;save R0 in R0 call stack position

                    rtrn


 

                    ;returns value in R0, R0=0 false, R0=1 true
stk-a-isempty:      ldi     r7 stk-a:hi
                    shl     r7 r7 8
                    addi    r7 stk-a:pt
                    ldi     r6 zero
                    ldw     r1 r7 r6             ;get stack-a top to R1

                    ldi     r0 one               ;R0=1 isEmpty, R0=0 not isEmpty
                    cmpi    r1 zero
                    ldi     r7 return:hi
                    shl     r7 r7 8
                    addi    r7 return:lo
                    bflag   r7 eq                   ;is empty, r0 = 1  
                    ldi     r0 zero              ;set R0 = 0, not is empty

return:             MVSR    r7
                    ldi     r6 R0IDX
                    stw     r0 r7 r6

                    rtrn
          
                    


