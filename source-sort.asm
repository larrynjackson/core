;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; insertion sort 
; O n2 worst case       ex: GFEDCBA
; O n  best case        ex: ABCDEFG
; O n2 average case     ex: DGFABCA
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;


ZERO:      DBYTE   0         ;constants ZERO and ONE
ONE:       DBYTE   1



buffer:lo:           dbyte  255
buffer:hi:           dbyte  66



start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
               
                     
                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo       
                    in      r0

                    ldi     r1 zero
                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo

bufferLength:       ldw     r2 r7 r1
                    addi    r1 one
                    cmpi    r2 zero
                    ldi     r5 bufferLength:hi
                    shl     r5 r5 8
                    addi    r5 bufferLength:lo    
                    bflag   r5 GT

                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo   
                    ;r0 is buffer pointer
                    ;r1 is buffer bufferLength
                    ;r2 is i
                    ;r3 is j 
                    ;r4 is temp
                    ;r5 is Array[j]
                    ldi     r6 zero     ;swaps counter
                    ldi     r2 zero     ;r2 is i
                    ldi     r3 one      ;r3 is j  
               
whilei<j:           ldw     r4 r0 r2    ;r4 is tmp
whilej<length:      ldw     r5 r0 r3    ;r5 is AR[j]
ifAR[j]<tmp:        cmpi    r5 zero
                    ldi     r7 endIf:hi
                    shl     r7 r7 8
                    addi    r7 endIf:lo
                    bflag   r7 EQ
                    cmp     r5 r4
                    ldi     r7 swap:hi
                    shl     r7 r7 8
                    addi    r7 swap:lo
                    bflag   r7 LT 
endIf:              ldi     r7 whilej<length:hi
                    shl     r7 r7 8
                    addi    r7 whilej<length:lo
                    addi    r3 one          ;j = j+1
                    cmp     r3 r1                
                    bflag   r7 LT           ;if j < len(buffer)
                    stw     r4 r0 r2        ;tmp to AR[i]
                    addi    r2 one          ;i = i+1
                    push    r2              ;save i
                    pop     r3              ;j = i
                    addi    r3 one          ;j = j+1
                    subi    r1 one          ;r1 = len(buffer)-1
                    cmp     r2 r1           ;i < len(buffer)-1
                    addi    r1 one          ;put r1 back to len(buffer)
                    ldi     r7 whilei<j:hi
                    shl     r7 r7 8
                    addi    r7 whilei<j:lo
                    bflag   r7 LT

                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    out     r7
                    halt

swap:               push    r4
                    push    r5
                    pop     r4
                    pop     r5
                    stw     r5 r0 r3    ;tmp to AR[j]
                    ldi     r7 endIf:hi
                    shl     r7 r7 8
                    addi    r7 endIf:lo
                    addi    r6 one
                    jump    r7

                     




