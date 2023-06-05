;some stuff that's not complete yet. work in progress

ZERO:      DBYTE   0         ;constants ZERO and ONE
ONE:       DBYTE   1
TWO:       DBYTE   2
THREE:     DBYTE   3
FOUR:      DBYTE   4
FIVE:      DBYTE   5

stk-a:lo:    dbyte   240
stk-a:hi:    dbyte   216

stk-b:lo:    dbyte   104
stk-b:hi:    dbyte   197

stk-c:lo:    dbyte   224
stk-c:hi:    dbyte   177

stk-d:lo:    dbyte   88
stk-d:hi:    dbyte   158

stk-e:lo:    dbyte   52
stk-e:hi:    dbyte   139

que-a:lo:    dbyte   72
que-a:hi:    dbyte   119

que-b:lo:    dbyte   192
que-b:hi:    dbyte   99

que-c:lo:    dbyte   56
que-c:hi:    dbyte   80

que-d:lo:    dbyte   176
que-d:hi:    dbyte   60



char*:     dbyte  42



buffer:lo:           dbyte  255
buffer:hi:           dbyte  66


start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0

bufferLength:       ldw     r2 r7 r5
                    addi    r5 one
                    cmpi    r2 zero
                    ldi     r6 bufferLength:hi
                    shl     r6 r6 8
                    addi    r6 bufferLength:lo    
                    bflag   r6 GT               ;r5 = len(buffer)
                    push    r5                  ;save buffer length
                 

                    ;read cmdline input into buffer
readCmdLine:        ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo       
                    in      r0
                   
                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo   
                    ldi     r5 zero
                    ldw     r3 r0 r5
                    cmpi    r3 char*
                    ldi     r7 endProgram:hi
                    shl     r7 r7 8
                    addi    r7 endProgram:lo
                    bflag   r7 EQ           

                    ldi     r5 zero
push-a:             ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo   
                    ldw     r3 r0 r5
                    push    r3                      ;save stack push value
                    ldi     r7 stk-a-push:hi
                    shl     r7 r7 8
                    addi    r7 stk-a-push:lo
                    call    r7
                    pop     r3                      ;remove pushed value from sys stack
                    ldi     r7 push-a:hi
                    shl     r7 r7 8
                    addi    r7 push-a:lo
                    pop     r4                      ;get buffer length
                    push    r4                      ;save buffer length
                    cmp     r5 r4                   ;is r5 less than len(buffer)
                    addi    r5 one                  ;step r5 to next buffer position
                    bflag   r7 LT

                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo
                    ldi     r1 zero                 ;r1 is buffer pointer

pop-a:              push    r2                      ;save space for isempty on stack
                    ldi     r7 stk-a-isempty:hi
                    shl     r7 r7 8
                    addi    r7 stk-a-isempty:lo
                    call    r7
                    pop     r2   
                    noop                   ;get isempty answer
                    cmpi    r2 one
                    ldi     r6 endMain:hi
                    shl     r6 r6 8
                    addi    r6 endMain:lo
                    bflag   r6 EQ


                    push    r2                      ;save space for pop value
                    ldi     r6 stk-a-pop:hi
                    shl     r6 r6 8
                    addi    r6 stk-a-pop:lo
                    call    r6
                    pop     r2                      ;get poped value
                    stw     r2 r0 r1
                    addi    r1 one
                    ldi     r2 zero
                    stw     r2 r0 r1
                    ldi     r6 pop-a:hi
                    shl     r6 r6 8
                    addi    r6 pop-a:lo
                    JUMP    r6



endMain:            out     r0
                  
                    ldi     r7 readCmdLine:hi
                    shl     r7 r7 8
                    addi    r7 readCmdLine:lo
                    jump    r7

 endProgram:        halt



stk-a-pop:          push    r0
                    push    r1
                    push    r2
                    push    r3

                    ldi     r0 stk-a:hi
                    shl     r0 r0 8
                    addi    r0 stk-a:lo
                    ldi     r1 zero
                    ldw     r2 r0 r1                ;r2 is stack pointer
                    subi    r2 one
                    stw     r2 r0 r1

                    ldw     r3 r0 r2
              
                    mvsr    r0
                    ldi     r1 five
                    stw     r3 r0 r1

                    pop     r3
                    pop     r2
                    pop     r1
                    pop     r0
                    rtrn


 stk-a-push:        push    r0
                    push    r1
                    push    r2
                    push    r3

                    MVSR    r0                
                    ldi     r1 five
                    ldw     r3 r0 r1            ;R3 = value to push
          
                    ldi     r0 stk-a:hi
                    shl     r0 r0 8
                    addi    r0 stk-a:lo
                    ldi     r1 zero
                    ldw     r2 r0 r1
                    addi    r2 one
                    stw     r3 r0 r2
                    stw     r2 r0 r1
                    pop     r3
                    pop     r2
                    pop     r1
                    pop     r0
                    rtrn


stk-a-isempty:      push    r0
                    push    r1
                    push    r2
                    push    r3

                    
          
                    ldi     r0 stk-a:hi
                    shl     r0 r0 8
                    addi    r0 stk-a:lo
                    ldi     r1 zero
                    ldw     r2 r0 r1                ;r2 is stack pointer
                    cmpi    r2 zero
                    ldi     r3 one
                    ldi     r0 return-isempty:hi
                    shl     r0 r0 8
                    addi    r0 return-isempty:lo
                    bflag   r0 EQ
                    ldi     r3 zero
return-isempty:     mvsr    r0
                    ldi     r1 five
                    stw     r3 r0 r1

                    pop     r3
                    pop     r2
                    pop     r1
                    pop     r0
                    rtrn



