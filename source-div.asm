


        
     

    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
    ;
    ;             SEAL
    ;  Simple Easy Assembly Language 
    ;                               
    ;  Integer division: dividend / divisor = quotient and remainder
    ;  SEAL is limited to immediate values and unsigned byte values
    ;  from 0 to 255. So to divide 385 by 8 some creative thinkig is
    ;  required. SEAL's registers and memory store 16 bit unsigned
    ;  integers. So to get numbers between 0 and 65535 into registers
    ;  we must convert the number to hexidecimal or binary then convert
    ;  the hi byte and lo byte to decimal then move them to a registers
    ;  in two moves. 385 in hex is 0x0181, binary 0000 0001 1000 0001
    ;  or hi byte = 1 and lo byte = 129. So we can load the hi byte, 
    ;  shift left 8 places then add to it the lo byte.
    ;
    ;  ldi  r0 1
    ;  shl  r0 r0 8
    ;  addi r0 129 
    ;
    ;  Now r0 stores the value 385. If the divisor is greater than 255,
    ;  then we would apply the same code.
    ;                               
    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;



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


    ;dividend:hi:     DBYTE       1
    ;dividend:lo:     DBYTE       129
    ;divisor:         DBYTE       8

    dividend:hi:     DBYTE       57
    dividend:lo:     DBYTE       145
    divisor:         DBYTE       247


                 
                    ; division subroutine

   
start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
        
        

                
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; set up call to divide ex: 385 / 8 = 48 remainder 1
                    ; 385 dividend
                    ; 8 divisor
                    ; push to the stack, dividend and divisor
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

                    ldi     r0 dividend:hi
                    shl     r0 r0 8
                    addi    r0 dividend:lo
                    ldi     r1 divisor

                    ldi     r7 divide:HI
                    shl     r7 r7 8
                    addi    r7 divide:LO
           
                    call    r7
          
                    
                    HALT


                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; stack r0 = dividend
                    ; stack r1 = divisor
                    ;
                    ; r4 = dividend
                    ; r5 = divisor
                    ;
                    ; r2 = quotient
                    ; r3 = remainder
                    ;
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    divide:         mvsr    r0
                    ldi     r1 r0idx
                    ldw     r4 r0 r1        ;r4 = dividend
                    ldi     r1 r1idx
                    ldw     r5 r0 r1        ;r5 = divisor
                    
                    ldi     r3 ZERO
                    add     r3 r3 r4  ; copy dividend to r3
                    ldi     r6 ZERO
    divloop:        sub     r4 r4 r5
                    addi    r6 ONE
                 
                    cmp     r3 r4
                    ldi     r7 divloop:HI
                    shl     r7 r7 8
                    addi    r7 divloop:LO
                    bflag   r7 gt
                    subi    r6 ONE
                    add     r4 r4 r5

                    ldi     r1 r2idx
                    stw     r6 r0 r1
                    ldi     r1 r3idx
                    stw     r4 r0 r1
                   
                    
                    RTRN


