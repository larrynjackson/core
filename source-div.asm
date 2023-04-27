


        
     

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



    ZERO:            DBYTE       0
    ONE:             DBYTE       1
    dividend:hi:     DBYTE       57
    dividend:lo:     DBYTE       145
    divisor:         DBYTE       247

                 
                    ; division subroutine

   
                    ldi     r7 setstack:HI
                    shl     r7 r7 8
                    addi    r7 setstack:LO
                    call    r7
        

                
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ; set up call to divide ex: 385 / 8 = 48 remainder 1
                    ; 385 dividend
                    ; 8 divisor
                    ; push to the stack, dividend and divisor
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

                    ldi     r4 dividend:hi
                    shl     r4 r4 8
                    addi    r4 dividend:lo
                    ldi     r5 divisor
                    push    r4
                    push    r5
                    ;noop

                    ldi     r7 divide:HI
                    shl     r7 r7 8
                    addi    r7 divide:LO
                    noop
                    call    r7

                    pop     r0  ;quotient
                    pop     r1  ;remainder
                   
                    
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
                    ; divide entry conditions uses r0 r3 r4 r5 r6 r7
                    ; caller should save registers if they contain active data
                    ; call should:
                    ;  - save any active registers on the stack
                    ;  - push dividend 
                    ;  - push divisor
                    ;  - return address on stack from call instruction
                    ;
                    ; The result will sit on top of the stack: quotient, remainder
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    divide:         pop     r0  ; hold return addr
                    pop     r5  ; divisor
                    pop     r4  ; dividend
                    noop
                    
                    ldi     r3 ZERO
                    add     r3 r3 r4  ; copy dividend to r3
                    ldi     r6 ZERO
    divloop:        sub     r4 r4 r5
                    addi    r6 ONE
                    noop
                    cmp     r3 r4
                    ldi     r7 divloop:HI
                    shl     r7 r7 8
                    addi    r7 divloop:LO
                    bflag   r7 gt
                    subi    r6 ONE
                    add     r4 r4 r5
                    push    r4         ; remainder
                    push    r6         ; quotient
                    push    r0         ; put back the return adr
                    
                    RTRN


