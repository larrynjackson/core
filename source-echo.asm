

;simple echo program. type * to quit.



ZERO:      DBYTE   0         ;constants ZERO and ONE
ONE:       DBYTE   1

char*:     dbyte  42



buffer:lo:           dbyte  255
buffer:hi:           dbyte  66



start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
                 

                    ;read cmdline input into buffer
readCmdLine:        ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo       
                    in      r0
                    noop
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
                    out     r0
                    noop
                    ldi     r7 readCmdLine:hi
                    shl     r7 r7 8
                    addi    r7 readCmdLine:lo
                    jump    r7

 endProgram:        ldi     r0 exitEcho:hi
                    shl     r0 r0 8
                    addi    r0 exitEcho:lo
                    out     r0
                    halt

 exitEcho:          dstring oh what a wonderful day


