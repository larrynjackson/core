


        
     

    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
    ;                               ;
    ;  comment intoduction header   ;
    ;                               ;
    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;




    multiplicand:    DBYTE       55
    multiplier:      DBYTE       7

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






                 
                    ; multiply subroutine

   
start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
        

                
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; set up call to multiply ex: 55 * 7
                    ; 55 multiplicand
                    ; 7  multiplier
                    ; set r0 = multiplicand
                    ; set r1 = multiplier
                    ; answer will be in r2 on return
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

                    ldi     r0 multiplicand      ;multiplicand
                    ldi     r1 multiplier        ;multiplier
                   
             

                    ldi     r7 multiply:HI
                    shl     r7 r7 8
                    addi    r7 multiply:LO
                 
                    call    r7
                    
                    HALT






                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ;
                    ; recover r4 and r5 from the call stack
                    ; r0 = multiplicand
                    ; r1 = multiplier
                    ;
                    ; r2 = answer
                    ; write the answer r2 back to r2 position on the call stack
                    ;  
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    multiply:       mvsr    r7
                    ldi     r6 r0idx
                    ldw     r0 r7 r6
                    ldi     r6 r1idx
                    ldw     r1 r7 r6

                  
                    
                    ldi     r2 zero
    mulloop:        add     r2 r2 r0
                    subi    r1 ONE
                  
                    cmpi    r1 ZERO
                    ldi     r5 mulloop:HI
                    shl     r5 r5 8
                    addi    r5 mulloop:LO
                    bflag   r5 ne
                    ldi     r6 r2idx
                    stw     r2 r7 r6
                    
                    RTRN


