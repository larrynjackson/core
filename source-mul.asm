


        
     

    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
    ;                               ;
    ;  comment intoduction header   ;
    ;                               ;
    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;



    ZERO:            DBYTE       0
    ONE:             DBYTE       1
    multiplicand:    DBYTE       55
    multiplier:      DBYTE       7

                 
                    ; multiply subroutine

   
                    ldi     r7 setstack:HI
                    shl     r7 r7 8
                    addi    r7 setstack:LO
                    call    r7
        

                
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; set up call to multiply ex: 55 * 7
                    ; 55 multiplicand
                    ; 7  multiplier
                    ; push to the stack, multiplicand then multiplier
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

                    ldi     r4 multiplicand      ;multiplicand
                    ldi     r5 multiplier        ;multiplier
                    push    r4
                    push    r5
                    ;noop

                    ldi     r7 multiply:HI
                    shl     r7 r7 8
                    addi    r7 multiply:LO
                    noop
                    call    r7

                    pop     r0
                    ;out     r0
                    
                    HALT


                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; set the current stack top to 0xFFFF (65535)
                    ; destroys r0 contents
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    setstack:       ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
                    rtrn



                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; multiply entry conditions uses r0 r4 r5 r6 r7
                    ; caller should save registers if they contain active data
                    ; call should:
                    ;  - save any active registers on the stack
                    ;  - push multiplicand 
                    ;  - push multiplier
                    ;  - return address on stack from call instruction
                    ;
                    ; The result will sit on top of the stack
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    multiply:       pop     r0  ; hold return addr
                    pop     r5  ; multiplier
                    pop     r4  ; multiplicand
                    noop
                    
                    ldi     r6 ZERO
    mulloop:        add     r6 r6 r4
                    subi    r5 ONE
                    noop
                    cmpi    r5 ZERO
                    ldi     r7 mulloop:HI
                    shl     r7 r7 8
                    addi    r7 mulloop:LO
                    bflag   r7 ne
                    push    r6         ; answer
                    push    r0         ; put back the return adr
                    
                    RTRN


