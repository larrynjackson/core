;used to test my new stack changes

ZERO:      DBYTE   0         
ONE:       DBYTE   1
TWO:       DBYTE   2
THREE:     DBYTE   3
FOUR:      DBYTE   4
FIVE:      DBYTE   5
SIX:       DBYTE   6
SEVEN:     DBYTE   7
EIGHT:     DBYTE   8

R0IDX:      DBYTE   2
R1IDX:      DBYTE   3
R2IDX:      DBYTE   4
R3IDX:      DBYTE   5
R4IDX:      DBYTE   6
R5IDX:      DBYTE   7
R6IDX:      DBYTE   8
R7IDX:      DBYTE   9




buffer:lo:           dbyte  255
buffer:hi:           dbyte  66

start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
                     
                    


main:               ldi     r0 5
                    ldi     r1 10
                    ldi     r2 15
                    ldi     r3 20
                    ldi     r4 25
                    ldi     r5 30
                    ldi     r6 35
                    ldi     r7 40
                    
                    ldi     r7 endMain:hi
                    shl     r7 r7 8
                    addi    r7 endMain:lo



                    call    r7                  ;call qs
                   
                    halt

endMain:            MVSR    r7
                    ldi     r6 r0idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r1idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6



                    ldi     r6 r2idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r3idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r4idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r5idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r6idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    ldi     r6 r7idx
                    ldw     r1 r7 r6
                    addi    r1 100
                    stw     r1 r7 r6

                    rtrn



