


        
     

        
        

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; simple 4 function calculator for integer arithmetic
; * / + -
; minimul error checking so beware of unrecognozed errors
; ex: largest positive 16 bit unsigned integer - 65535
; operations may exceed that value and rollover
; multiplication is achieved by addition
; division is achieved by subtraction
; so multiply and divide operations on large numbers will take a very long time
; subtraction may show a negative number, but negative numbers are not legal
; division produces quotient and remainder. ex: 56/5 yields 11 r 1
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;




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


chara:     dbyte  65         ;ascii chars A-Z
charb:     dbyte  66
charc:     dbyte  67
chard:     dbyte  68
chare:     dbyte  69
charf:     dbyte  70
charg:     dbyte  71
charh:     dbyte  72
chari:     dbyte  73
charj:     dbyte  74
chark:     dbyte  75
charl:     dbyte  76
charm:     dbyte  77
charn:     dbyte  78
charo:     dbyte  79
charp:     dbyte  80
charq:     dbyte  81
charr:     dbyte  82
chars:     dbyte  83
chart:     dbyte  84
charu:     dbyte  85
charv:     dbyte  86
charw:     dbyte  87
charx:     dbyte  88
chary:     dbyte  89
charz:     dbyte  90
charsp:    dbyte  32

char+:     dbyte  43           ;ascii operators +,-,*,/ and =
char-:     dbyte  45
char*:     dbyte  42
char/:     dbyte  47
char=:     dbyte  61

num0:      dbyte     48        ;ascii digits 0-9
num1:      dbyte     49
num2:      dbyte     50
num3:      dbyte     51
num4:      dbyte     52
num5:      dbyte     53
num6:      dbyte     54
num7:      dbyte     55
num8:      dbyte     56
num9:      dbyte     57

operator:lo:         dbyte  250
operator:hi:         dbyte  64

opdLvalue:lo:        dbyte  253
opdLvalue:hi:        dbyte  64

opdLtop:lo:          dbyte  254
opdLstk:lo:          dbyte  255
opdLstk:hi:          dbyte  64

opdRvalue:lo:        dbyte  253
opdRvalue:hi:        dbyte  65

opdRtop:lo:          dbyte  254
opdRstk:lo:          dbyte  255
opdRstk:hi:          dbyte  65

bufferNext:lo:       dbyte  254
buffer:lo:           dbyte  255
buffer:hi:           dbyte  66



start:               ldi     r0 ZERO
                     not     r0 r0
                     ldsr    r0
               
                     ;initialize Left digit stack, 
                     ;Right digit stack and operator storage
readInput:           ldi    r7 initopcLRstk:HI
                     shl    r7 r7 8
                     addi   r7 initopcLRstk:LO 
                     call   r7

                     ;initialize input/output buffer
                     ldi    r7 initBuffer:hi
                     shl    r7 r7 8
                     addi   r7 initBuffer:lo
                     call   r7
                     ldi    r6 zero
                     ldi    r7 putR6Operator:hi
                     shl    r7 r7 8
                     addi   r7 putR6Operator:lo
                     call   r7

                     ;read cmdline input into buffer
                     ldi    r0 buffer:hi
                     shl    r0 r0 8
                     addi   r0 buffer:lo       
                     in     r0

                     ldi    r7 seeR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 seeR6BufferNext:lo
                     call   r7
                     ldi    r7 endProgram:hi
                     shl    r7 r7 8
                     addi   r7 endProgram:lo
                     cmpi   r6 charq
                     bflag  r7 EQ


                     ; GET Left side digits into left operand stack
                     ; GET operator into operator storage
                     ldi    r7 readopdL:hi
                     shl    r7 r7 8
                     addi   r7 readopdL:lo
                     call   r7

                     ; GET Right side digits into right operand stack
                     ldi    r7 readopdR:hi
                     shl    r7 r7 8
                     addi   r7 readopdR:lo
                     call   r7

                     ;validate input command ex: 45 + 34
                     ldi    r7 testInput:hi
                     shl    r7 r7 8
                     addi   r7 testInput:lo
                     call   r7


                     ;convert left operand char digits into a number
                     ldi    r7 convertOpdLstk:hi
                     shl    r7 r7 8
                     addi   r7 convertOpdLstk:lo
                     call   r7

                     ;convert right operand char digits into a number
                     ldi    r7 convertOpdRstk:hi
                     shl    r7 r7 8
                     addi   r7 convertOpdRstk:lo
                     call   r7

                     ;evaluate converted input (left number operator right number)
                     ldi    r7 evaluate:hi
                     shl    r7 r7 8
                     addi   r7 evaluate:lo
                     call   r7

                     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 buffer:lo

                     out    r7
                     ldi    r7 readInput:hi
                     shl    r7 r7 8
                     addi   r7 readInput:lo
                     jump   r7

 endProgram:         halt


callInputError:      ldi    r7 error:hi
                     shl    r7 r7 8
                     addi   r7 error:lo
                     call   r7
                     ldi    r7 readInput:hi
                     shl    r7 r7 8
                     addi   r7 readInput:lo
                     jump   r7
                     


testInput:           ldi    r7 opdLstk:hi
                     shl    r7 r7 8
                     addi   r7 opdLtop:lo
                     ldi    r0 zero
                     ldw    r1 r7 r0
  
                     cmpi   r1 1
                     ldi    r5 callInputError:hi
                     shl    r5 r5 8
                     addi   r5 callInputError:lo
                     bflag  r5 EQ

                     ldi    r7 opdRstk:hi
                     shl    r7 r7 8
                     addi   r7 opdRtop:lo
                     ldi    r0 zero
                     ldw    r1 r7 r0

                     cmpi   r1 1
                     bflag  r5 EQ

                     ldi    r6 testInputOk:hi
                     shl    r6 r6 8
                     addi   r6 testInputOk:lo
                     ldi    r7 operator:hi
                     shl    r7 r7 8
                     addi   r7 operator:lo
                     ldi    r0 zero
                     ldw    r1 r7 r0
                     cmpi   r1 char*                
                     bflag  r6 EQ
                     cmpi   r1 char+               
                     bflag  r6 EQ
                     cmpi   r1 char-                
                     bflag  r6 EQ
                     cmpi   r1 char/                 
                     bflag  r6 EQ

                     ldi    r5 callInputError:hi
                     shl    r5 r5 8
                     addi   r5 callInputError:lo
                     jump   r5


testInputOk:         rtrn







                     ;call function. reads and places left side digits onto opdLeft stack.
                     ;reads and places operator in save operator location
                     ;read next character from the buffer into r6
                     
readopdL:            ldi    r7 getR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 getR6BufferNext:lo
                     call   r7
                     ;on return the next character is in R6

                     ;eat beginning spaces
                     cmpi   r6 charsp
                     ldi    r7 readopdL:hi
                     shl    r7 r7 8
                     addi   r7 readopdL:lo
                     bflag  r7 EQ

                   

                     ;look is R6 digit. R1 0(false), 1(true)
                     ldi    r7 isR6Digit:hi
                     shl    r7 r7 8
                     addi   r7 isR6Digit:lo
                     call   r7
                     ;on return R1 contains 0=false or 1=true
                 

                     ldi    r7 callPushR6OpdLeft:hi
                     shl    r7 r7 8
                     addi   r7 callPushR6OpdLeft:lo
                     cmpi   r1 one
                     bflag  r7 EQ

                     cmpi   r6 charsp
                     ldi    r7 readopdL:hi
                     shl    r7 r7 8
                     addi   r7 readopdL:lo
                     bflag  r7 EQ

                     

                     ;is operator

                     ldi    r7 callPutR6Operator:hi
                     shl    r7 r7 8
                     addi   r7 callPutR6Operator:lo

                     cmpi   r6 char+
                     bflag  r7 EQ
                     cmpi   r6 char-
                     bflag  r7 EQ
                     cmpi   r6 char*
                     bflag  r7 EQ
                     cmpi   r6 char/
                     bflag  r7 EQ

                     rtrn


 callPutR6Operator:  ldi    r7 putR6Operator:hi
                     shl    r7 r7 8
                     addi   r7 putR6Operator:lo
                     call   r7          
                     rtrn               






callPushR6OpdRight:  ldi    r7 pushR6OpdRight:hi
                     shl    r7 r7 8
                     addi   r7 pushR6OpdRight:lo
                     call   r7
                     ldi    r7 readopdR:hi
                     shl    r7 r7 8
                     addi   r7 readopdR:lo
                     jump   r7
                   

                     ;read next character from the buffer into r6
                     
readopdR:            ldi    r7 getR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 getR6BufferNext:lo

                     call   r7
                     ;on return next buffer element is in R6
            

                     ;eat beginning spaces
                     cmpi   r6 charsp
                     ldi    r7 readopdR:hi
                     shl    r7 r7 8
                     addi   r7 readopdR:lo
                     bflag  r7 EQ

                   

                     ;look is R6 digit. R1 0(false), 1(true)
                     ldi    r7 isR6Digit:hi
                     shl    r7 r7 8
                     addi   r7 isR6Digit:lo
                     call   r7
                     ;on return R1=0 (false), R1=1 (true)
            

                     ldi    r7 callPushR6OpdRight:hi
                     shl    r7 r7 8
                     addi   r7 callPushR6OpdRight:lo
                     cmpi   r1 one
                
                     bflag  r7 EQ

                     cmpi   r6 charsp
                     ldi    r7 readopdR:hi
                     shl    r7 r7 8
                     addi   r7 readopdR:lo
                     bflag  r7 EQ

                     rtrn


callEvalError:       ldi    r7 error:hi
                     shl    r7 r7 8
                     addi   r7 error:lo
                     call   r7

                     halt

evaluate:            ldi    r7 initBuffer:hi
                     shl    r7 r7 8
                     addi   r7 initBuffer:lo
                     call   r7
                     ldi    r7 operator:hi
                     shl    r7 r7 8
                     addi   r7 operator:lo
                     ldi    r0 zero
                     ldw    r3 r7 r0     ;operator

                     ldi    r7 domul:hi
                     shl    r7 r7 8
                     addi   r7 domul:lo                  
                     cmpi   r3 char*                
                     bflag  r7 EQ
                     ldi    r7 doadd:hi
                     shl    r7 r7 8
                     addi   r7 doadd:lo
                     cmpi   r3 char+               
                     bflag  r7 EQ
                     ldi    r7 dosub:hi
                     shl    r7 r7 8
                     addi   r7 dosub:lo
                     cmpi   r3 char-                
                     bflag  r7 EQ
                     ldi    r7 dodiv:hi
                     shl    r7 r7 8
                     addi   r7 dodiv:lo
                     cmpi   r3 char/                 
                     bflag  r7 EQ

                     ldi    r5 callEvalError:hi
                     shl    r5 r5 8
                     addi   r5 callEvalError:lo
                     call   r5



doadd:               ldi    r7 opdLvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdLvalue:lo
                     ldi    r0 zero
                     ldw    r2 r7 r0

                     ldi    r7 opdRvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdRvalue:lo
                     ldi    r0 zero
                     ldw    r4 r7 r0

                     add    r1 r2 r4      ;answer

                     ldi    r7 convAns2Str:hi
                     shl    r7 r7 8
                     addi   r7 convAns2Str:lo
                     call   r7
                     rtrn


ifIsSubNegative:     sub    r1 r4 r2      ;answer in r1
                     ldi    r6 char-      ;sign
                     ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo

                     call   r7

                     ldi    r7 ifIsNotSubNegative:hi
                     shl    r7 r7 8
                     addi   r7 ifIsNotSubNegative:lo 
                     jump   r7 
                                       
     
dosub:               ldi    r7 opdLvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdLvalue:lo
                     ldi    r0 zero
                     ldw    r2 r7 r0

                     ldi    r7 opdRvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdRvalue:lo
                     ldi    r0 zero
                     ldw    r4 r7 r0

                     ldi    r7 ifIsSubNegative:hi
                     shl    r7 r7 8
                     addi   r7 ifIsSubNegative:lo         
                     cmp    r2 r4
                     bflag  r7 LT
                     sub    r1 r2 r4      ;answer in r1
ifIsNotSubNegative:  ldi    r7 convAns2Str:hi
                     shl    r7 r7 8
                     addi   r7 convAns2Str:lo
                     call   r7
                     rtrn

domul:               ldi    r7 opdLvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdLvalue:lo
                     ldi    r0 zero
                     ldw    r2 r7 r0

                     ldi    r7 opdRvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdRvalue:lo
                     ldi    r0 zero
                     ldw    r4 r7 r0


                     push   r2          ;multiplcand
                     pop    r3
                     push   r4          ;multiplier
                     pop    r6

                     ldi    r7 multiply:hi
                     shl    r7 r7 8
                     addi   r7 multiply:lo
                     call   r7
                   ;on return answer is in R1
                     ldi    r7 convAns2Str:hi
                     shl    r7 r7 8
                     addi   r7 convAns2Str:lo
                     call   r7
                     rtrn



                    
dodiv:               ldi    r7 opdLvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdLvalue:lo
                     ldi    r0 zero
                     ldw    r2 r7 r0

                     ldi    r7 opdRvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdRvalue:lo
                     ldi    r0 zero
                     ldw    r4 r7 r0

                     push   r2            ;dividend
                     pop    r0
                     push   r4            ;divisor
                     pop    r1
                     ldi    r7 divide:hi
                     shl    r7 r7 8
                     addi   r7 divide:lo
                     call   r7
                     push   r2
                     pop    r1            ;quotient answer
                    
                     ldi    r7 convAns2Str:hi
                     shl    r7 r7 8
                     addi   r7 convAns2Str:lo
                     call   r7
                     ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo

                     ldi    r6 charsp                  
                     call   r7

                     ldi    r6 charr                
                     call   r7
                
                     ldi    r6 charsp                      
                     call   r7
                    
                     push    r3
                     pop    r1                   ;remainder from initial divide
                     ldi    r7 convAns2Str:hi
                     shl    r7 r7 8
                     addi   r7 convAns2Str:lo
                     call   r7


                     rtrn

                     ;r1 contains answer number
convAns2Str:         push   r1            ;dividend
                     pop    r0

                     ldi    r4 39
                     shl    r4 r4 8
                     addi   r4 16
                     push   r4            ;divisor
                     pop    r1

                     ldi    r7 divide:hi
                     shl    r7 r7 8
                     addi   r7 divide:lo
                     call   r7

                     push   r2
                     pop    r6
                     push   r3
                     pop    r2
                     
                     ldi    r7 write10k:hi
                     shl    r7 r7 8
                     addi   r7 write10k:lo
                     cmpi   r6 zero
                     bflag  r7 GT
done10kwrite:        push   r2
                     pop    r0

                     ldi    r4 3
                     shl    r4 r4 8
                     addi   r4 232
                     push   r4            ;divisor    
                     pop    r1

                     ldi    r7 divide:hi
                     shl    r7 r7 8
                     addi   r7 divide:lo
                     call   r7

                     push   r2
                     pop    r6
                     push   r3
                     pop    r2

                     ldi    r7 write1k:hi
                     shl    r7 r7 8
                     addi   r7 write1k:lo
                     cmpi   r6 zero
                     bflag  r7 GT
done1kwrite:         push   r2
                     pop    r0

                     ldi    r4 100
                     push   r4            ;divisor   
                     pop    r1

                     ldi    r7 divide:hi
                     shl    r7 r7 8
                     addi   r7 divide:lo
                     call   r7


                     push   r2
                     pop    r6
                     push   r3
                     pop    r2

                     ldi    r7 write100:hi
                     shl    r7 r7 8
                     addi   r7 write100:lo
                     cmpi   r6 zero
                     bflag  r7 GT
done100write:        push   r2
                     pop    r0

                     ldi    r4 10
                     push   r4            ;divisor
                     pop    r1

                     ldi    r7 divide:hi
                     shl    r7 r7 8
                     addi   r7 divide:lo
                     call   r7


                     push   r2
                     pop    r6
                     push   r3
                     pop    r2

                     ldi    r7 write10:hi
                     shl    r7 r7 8
                     addi   r7 write10:lo
                     cmpi   r6 zero
                     bflag  r7 GT
done10write:         push   r2
                     pop    r6

                     ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     addi   r6 num0
                     call   r7
                     rtrn





write10k:            ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     addi   r6 num0
                     call   r7
                     ldi    r7 done10kwrite:hi
                     shl    r7 r7 8
                     addi   r7 done10kwrite:lo
                     jump   r7

write1k:             ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     addi   r6 num0
                     call   r7
                     ldi    r7 done1kwrite:hi
                     shl    r7 r7 8
                     addi   r7 done1kwrite:lo
                     jump   r7

write100:            ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     addi   r6 num0
                     call   r7
                     ldi    r7 done100write:hi
                     shl    r7 r7 8
                     addi   r7 done100write:lo
                     jump   r7

write10:             ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     addi   r6 num0
                     call   r7
                     ldi    r7 done10write:hi
                     shl    r7 r7 8
                     addi   r7 done10write:lo
                     jump   r7





                       
error:               ldi    r7 initBuffer:hi
                     shl    r7 r7 8
                     addi   r7 initBuffer:lo
                     call   r7

                     ldi    r7 putR6BufferNext:hi
                     shl    r7 r7 8
                     addi   r7 putR6BufferNext:lo
                     ldi    r6 chare  
                     call   r7
                     ldi    r6 charr 
                     call   r7
                     ldi    r6 charr 
                     call   r7
                     ldi    r6 charo
                     call   r7
                     ldi    r6 charr
                     call   r7
                     ldi    r6 zero
                     call   r7

                     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 buffer:lo
                     out    r7
                     
                     rtrn






                     ;checks for digit character in R6. Returns 0(false), 1(true) in R1
returnIsDigitTrue:   ldi r1 one
                     mvsr   r5
                     ldi    r4 r1idx
                     stw    r1 r5 r4

                     rtrn 

 isR6Digit:          mvsr   r5
                     ldi    r4 r6idx
                     ldw    r6 r5 r4
 
                     ldi    r7 returnIsDigitTrue:hi
                     shl    r7 r7 8
                     addi   r7 returnIsDigitTrue:lo
                     cmpi   r6 num0
                     bflag  r7 EQ
                     cmpi   r6 num1
                     bflag  r7 EQ
                     cmpi   r6 num2
                     bflag  r7 EQ
                     cmpi   r6 num3
                     bflag  r7 EQ
                     cmpi   r6 num4
                     bflag  r7 EQ
                     cmpi   r6 num5
                     bflag  r7 EQ
                     cmpi   r6 num6
                     bflag  r7 EQ
                     cmpi   r6 num7
                     bflag  r7 EQ
                     cmpi   r6 num8
                     bflag  r7 EQ
                     cmpi   r6 num9
                     bflag  r7 EQ
                     ldi    r1 zero

                     mvsr   r5
                     ldi    r4 r1idx
                     stw    r1 r5 r4

                     RTRN





                    ;branch function multiply r3 * r6 
callLMultiply:       ldi    r7 multiply:hi
                     shl    r7 r7 8
                     addi   r7 multiply:lo
          
                     call   r7
                   
                     add    r2 r1 r2
                     ldi    r7 opdLvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdLvalue:lo
                     ldi    r0 zero
                     stw    r2 r7 r0

                     ldi    r7 convertLloop:hi
                     shl    r7 r7 8
                     addi   r7 convertLloop:lo
                     jump   r7


convertComplete:     rtrn


                     ;call function. 
                     ;new conversion code. digit to convert is R6
                     ;storing converted digit in R2
convertOpdLstk:      ldi    r5 zero
                     ldi    r2 zero                    
convertLloop:        ldi    r7 popR6OpdLeft:hi
                     shl    r7 r7 8
                     addi   r7 popR6OpdLeft:lo

                     call   r7
                     ;on return popped digit is in R6

                     ldi    r7 convertComplete:hi
                     shl    r7 r7 8
                     addi   r7 convertComplete:lo
                     cmpi   r6 zero
                     bflag  r7 EQ

                     addi   r5 one  
                     subi   r6 num0 
                     
                     ldi    r7 callLMultiply:hi
                     shl    r7 r7 8
                     addi   r7 callLMultiply:lo

                     ldi    r3 ONE 
                     cmpi   r5 ONE
                     bflag  r7 EQ
                     ;if r5 is 1, then multiply 1 * r6 and add to r2

                     ldi    r3 10
                     cmpi   r5 TWO
                     bflag  r7 EQ
                     ;if r5 is 2, then multiply 10 * r6 and add to r2

                     ldi    r3 100
                     cmpi   r5 THREE
                     bflag  r7 EQ
                     ;if r5 is 3, then multiply 100 * r6 and add to r2
                  
                     ldi    r3 3
                     shl    r3 r3 8
                     addi   r3 232
                     cmpi   r5 FOUR
                     bflag  r7 EQ
                     ;if r5 is 4, then multiply 1000 * r6 and add to r2

                     ldi    r3 39
                     shl    r3 r3 8
                     addi   r3 16
                     cmpi   r5 FIVE
                     bflag  r7 EQ
                     ;if r5 is 5, then multiply 10000 * r6 and add to r2

                     ldi    r7 error:hi
                     shl    r7 r7 8
                     addi   r7 error:lo
                     call   r7
                     rtrn



                    ;branch function. r3 * r6. 
callRMultiply:       ldi    r7 multiply:hi
                     shl    r7 r7 8
                     addi   r7 multiply:lo
                     call   r7
                     add    r4 r1 r4
                     ldi    r7 opdRvalue:hi
                     shl    r7 r7 8
                     addi   r7 opdRvalue:lo
                     ldi    r0 zero
                     stw    r4 r7 r0

                     ldi    r7 convertRloop:hi
                     shl    r7 r7 8
                     addi   r7 convertRloop:lo
                     jump   r7





                     ;new conversion code. digit to convert is R6
                     ;storing converted digit in R4
convertOpdRstk:      ldi    r5 zero
                     ldi    r4 zero                    
convertRloop:        ldi    r7 popR6OpdRight:hi
                     shl    r7 r7 8
                     addi   r7 popR6OpdRight:lo
                     call   r7
                     ;on return R6=popped value
                   

                     ldi    r7 convertComplete:hi
                     shl    r7 r7 8
                     addi   r7 convertComplete:lo
                     cmpi   r6 zero
                     bflag  r7 EQ

                     addi   r5 one  

                     
                     subi   r6 num0 
                     
                     ldi    r7 callRMultiply:hi
                     shl    r7 r7 8
                     addi   r7 callRMultiply:lo

                     ldi    r3 ONE 
                     cmpi   r5 ONE
                     bflag  r7 EQ
                     ;if r5 is 1, then multiply 1 * r6 and add to r4

                     ldi    r3 10
                     cmpi   r5 TWO
                     bflag  r7 EQ
                     ;if r5 is 2, then multiply 10 * r6 and add to r4

                     ldi    r3 100
                     cmpi   r5 THREE
                     bflag  r7 EQ
                     ;if r5 is 3, then multiply 100 * r6 and add to r4
                     
                     ldi    r3 3
                     shl    r3 r3 8
                     addi   r3 232
                     cmpi   r5 FOUR
                     bflag  r7 EQ
                     ;if r5 is 4, then multiply 1000 * r6 and add to r4

                     ldi    r3 39
                     shl    r3 r3 8
                     addi   r3 16
                     cmpi   r5 FIVE
                     bflag  r7 EQ
                     ;if r5 is 5, then multiply 10000 * r6 and add to r4

                     ldi    r7 error:hi
                     shl    r7 r7 8
                     addi   r7 error:lo
                     call   r7
                     rtrn


                     ; call function. places next buffer Character in R6.
                     ; does not remove character from the buffer
seeR6BufferNext:     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 bufferNext:lo
                     ldi    r1 zero
                     ldw    r0 r7 r1      ;get buffer next index
                     ldi    r5 buffer:hi
                     shl    r5 r5 8
                     addi   r5 buffer:lo
                     ldw    r6 r5 r0

                     mvsr   r0
                     ldi    r1 r6idx
                     stw    r6 r0 r1
                     rtrn



                     ; call function. places next buffer Character in R6.
                     ; does remove character from the buffer
getR6BufferNext:     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 bufferNext:lo
                     ldi    r1 zero
                     ldw    r0 r7 r1      ;get buffer next index
                     ldi    r5 buffer:hi
                     shl    r5 r5 8
                     addi   r5 buffer:lo
                     ldw    r6 r5 r0
                     addi   r0 one
                     stw    r0 r7 r1

                     mvsr   r0
                     ldi    r1 r6idx
                     stw    r6 r0 r1
                     rtrn


                     ; call function. Writes R6 character to the buffer
putR6BufferNext:     mvsr   r0
                     ldi    r1 r6idx
                     ldw    r6 r0 r1

                     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 bufferNext:lo
                     ldi    r1 zero
                     ldw    r0 r7 r1      ;get buffer next index

                     ldi    r5 buffer:hi
                     shl    r5 r5 8
                     addi   r5 buffer:lo
                     stw    r6 r5 r0      ;store R6 at buffer[next]
                     addi   r0 one
                     stw    r1 r5 r0      ;store 0 (r1) at buffer[next] EOS
                     stw    r0 r7 r1      ;save next buffer location
                     rtrn



callPushR6OpdLeft:   ldi    r7 pushR6OpdLeft:hi
                     shl    r7 r7 8
                     addi   r7 pushR6OpdLeft:lo
                     call   r7

                     ldi    r7 readopdL:hi
                     shl    r7 r7 8
                     addi   r7 readopdL:lo
                     jump   r7


                     ;call function. pushes input R6 onto OpdLeft stack
pushR6OpdLeft:       mvsr   r7
                     ldi    r5 r6idx
                     ldw    r6 r7 r5

                     ldi    r7 opdLstk:Hi
                     shl    r7 r7 8
                     addi   r7 opdLtop:LO
                     ldi    r1 zero
                     ldw    r0 r7 r1       ;get top index
                     ldi    r5 opdLstk:Hi
                     shl    r5 r5 8
                     addi   r5 opdLstk:LO
                     stw    r6 r5 r0
                     addi   r0 one
                     stw    r0 r7 r1
                     rtrn

emptyStack:          ldi    r6 zero
                     mvsr   r7
                     ldi    r5 r6idx
                     stw    r6 r7 r5
                     rtrn

                     ;call function. pops opdLeft stack. places into R6 for return.
                     ;return digit in R6
popR6OpdLeft:        ldi    r7 opdLstk:Hi
                     shl    r7 r7 8
                     addi   r7 opdLtop:LO
                     ldi    r1 zero
                     ldw    r0 r7 r1       ;get top index
                     subi   r0 one
                     ldi    r3 emptyStack:hi
                     shl    r3 r3 8
                     addi   r3 emptyStack:lo
                     cmpi   r0 zero
                     bflag  r3 EQ
                     ldi    r5 opdLstk:Hi
                     shl    r5 r5 8
                     addi   r5 opdLstk:LO
                     ldw    r6 r5 r0
                     stw    r0 r7 r1

                     mvsr   r7
                     ldi    r5 r6idx
                     stw    r6 r7 r5

                     rtrn                     
                     
                     ;call function. pushes input R6 onto OpdRight stack
pushR6OpdRight:      mvsr   r7
                     ldi    r5 r6idx
                     ldw    r6 r7 r5

                     ldi    r7 opdRstk:Hi
                     shl    r7 r7 8
                     addi   r7 opdRtop:LO
                     ldi    r1 zero
                     ldw    r0 r7 r1       ;get top index
                     ldi    r5 opdRstk:Hi
                     shl    r5 r5 8
                     addi   r5 opdRstk:LO
                     stw    r6 r5 r0
                     addi   r0 one
                     stw    r0 r7 r1
                     rtrn

                     ;call function. pops opdRight stack. places into R6 for return.
                     ;return digit in R6
popR6OpdRight:       ldi    r7 opdRstk:Hi
                     shl    r7 r7 8
                     addi   r7 opdRtop:LO
                     ldi    r1 zero
                     ldw    r0 r7 r1       ;get top index
                     subi   r0 one
                     ldi    r3 emptyStack:hi
                     shl    r3 r3 8
                     addi   r3 emptyStack:lo
                     cmpi   r0 zero
                     bflag  r3 EQ
                     ldi    r5 opdRstk:Hi
                     shl    r5 r5 8
                     addi   r5 opdRstk:LO
                     ldw    r6 r5 r0
                     stw    r0 r7 r1

                     mvsr   r7
                     ldi    r5 r6idx
                     stw    r6 r7 r5

                     rtrn                     





                    ;call function. saves the operator in R6
putR6Operator:       mvsr   r0
                     ldi    r1 r6idx
                     ldw    r6 r0 r1

                     ldi    r7 operator:Hi
                     shl    r7 r7 8
                     addi   r7 operator:LO
                     ldi    r1 zero
                     stw    r6 r7 r1
                     rtrn


                    ;call function. returns the stored operator into R6.
getR6Operator:       ldi    r7 operator:Hi
                     shl    r7 r7 8
                     addi   r7 operator:LO
                     ldi    r1 zero
                     ldw    r6 r7 r1

                     mvsr   r7
                     ldi    r5 r6idx
                     stw    r6 r7 r5

                     rtrn
             


                    ;call function. no input. no return
initopcLRstk:        ldi     r1 zero
                     ldi     r0 one
                     ldi     r7 opdLstk:HI  
                     shl     r7 r7 8
                     addi    r7 opdLtop:LO  
                     stw     r0 r7 r1        ;set opdLtop 1

                     ldi     r7 opdRstk:HI  
                     shl     r7 r7 8
                     addi    r7 opdRtop:LO  
                     stw     r0 r7 r1        ;set opdRtop 1
                     rtrn




                    ;call function. no input. no return
initBuffer:          ldi    r0 zero
                     ldi    r7 buffer:hi
                     shl    r7 r7 8
                     addi   r7 bufferNext:lo
                     stw    r0 r7 r0
                     rtrn


                    

                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
                    ;
                    ; recover r4 and r5 from the call stack
                    ; r0 = multiplicand
                    ; r1 = multiplier
                    ;
                    ; r2 = answermultiply
                    ; write the answer r2 back to r2 position on the call stack
                    ;  
                    ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

    multiply:       mvsr    r7
                    ldi     r6 r3idx
                    ldw     r0 r7 r6
                    ldi     r6 r6idx
                    ldw     r1 r7 r6

                  
                    
                    ldi     r2 zero
    mulloop:        add     r2 r2 r0
                    subi    r1 ONE
                  
                    cmpi    r1 ZERO
                    ldi     r5 mulloop:HI
                    shl     r5 r5 8
                    addi    r5 mulloop:LO
                    bflag   r5 ne

                    mvsr    r7
                    ldi     r6 r1idx
                    stw     r2 r7 r6
                    
                    RTRN










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


