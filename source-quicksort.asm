;old quicksort. not working with new code. working on it.

ZERO:      DBYTE   0         
ONE:       DBYTE   1
TWO:       DBYTE   2
THREE:     DBYTE   3
FOUR:      DBYTE   4



buffer:lo:           dbyte  255
buffer:hi:           dbyte  66

start:              ldi     r0 ZERO
                    not     r0 r0
                    ldsr    r0
                     
                    ldi     r0 buffer:hi
                    shl     r0 r0 8
                    addi    r0 buffer:lo       
                    in      r0

                    ldi     r5 zero
                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo

bufferLength:       ldw     r2 r7 r5
                    addi    r5 one
                    cmpi    r2 zero
                    ldi     r6 bufferLength:hi
                    shl     r6 r6 8
                    addi    r6 bufferLength:lo    
                    bflag   r6 GT               ;r5 = len(buffer)
                    subi    R5 ONE              ;r5 is now size
                    subi    r5 one              ;r5 is now right (size-1)
                

main:               ldi     r7 qs:hi
                    shl     r7 r7 8
                    addi    r7 qs:lo
                    ldi     r4 zero             ;r4 = 0

                    ;create a consistant stack frame where 
                    ;SR + 1 is the return address from call
                    ;SR + 2 is R
                    ;SR + 3 is L
                    ;SR + 4 is I 
                    
                    push    r4                  ;push i=0 for consistant stack   
                    push    r4                  ;push L on the stack
                    push    r5                  ;push R on the stack
                    call    r7                  ;call qs
                   

endMain:            ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    out     r7
        
                    halt



qs:                 MVSR    r0                
                    ldi     r1 two
                    ldw     r5 r0 r1            ;load R5 with R
                    ldi     r1 three
                    ldw     r4 r0 r1            ;load r4 with L  

                    ;check for a negative rollover to a large positive
                    ;number. ex: R should be -1 but rolls into 65535
                    ;if caught jump to endif otherwise set R to zero
                    addi    r5 one
                    cmpi    r5 zero
                    ldi     r7 endIf:hi
                    shl     r7 r7 8
                    addi    r7 endIf:lo
                    bflag   r7 EQ

                    subi    r5 one

                    cmp     r5 r4               ; if r > l do if
             
                    ldi     r7 doIf:hi          ;if l > r or l == r jump doIf
                    shl     r7 r7 8
                    addi    r7 doIf:lo
                    bflag   r7 GT

endIf:              pop     r6
                    pop     r0
                    pop     r0
                    pop     r0
                    push    r6  
                     
                    rtrn

doIf:               ldi     r7 partition:hi
                    shl     r7 r7 8
                    addi    r7 partition:lo
                    jump    r7


                
retPartition:       MVSR    r0
                    ldi     r1 four
                    ldw     r2 r0 r1            ;r2 = i
                    subi    r2 one              ;r2 = i-1
                    ldi     r1 three
                    stw     r4 r0 r1            ;r4 = L

                    push    r2                  ;push dummy i value
                    push    r4                  ;push r4 L
                    push    r2                  ;push i-1 R

                    ldi     r7 qs:hi
                    shl     r7 r7 8
                    addi    r7 qs:lo   
                    call    r7
      
                    MVSR    r0
                    ldi     r1 four
                    ldw     r2 r0 r1            ;r2 = i
                    push    r2                  ;put back i on the stack
                    addi    r2 one
                    push    r2                  ;push L (i+1)
                    ldi     r1 two
                    ldw     r5 r0 r1
                    push    r5                  ; push R

                   
                    ldi     r7 qs:hi
                    shl     r7 r7 8
                    addi    r7 qs:lo
                    call    r7

                    pop     r6
                    pop     r0
                    pop     r0
                    pop     r0
                    push    r6
                    rtrn


                    
partition:          MVSR    r0
                    ldi     r1 two
                    ldw     r5 r0 r1                ;r5 = R
                    ldi     r1 three
                    ldw     r4 r0 r1                ;r4 = L

                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    ldw     r0 r7 r5                ;v = a[r]
                    push    r4                      ;L
                    pop     r2                      ;i = L
                    subi    r2 one                  ;i = l-1
                    push    r5                      ;r
                    pop     r3                      ;j = R

doForOne:           addi    r2 one                  ;i++

                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    ldw     r6 r7 r2                ;a[i] into r6
                    cmp     r6 r0                   ;if r6 a[i] >= v

                    ldi     r7 doForTwo:hi
                    shl     r7 r7 8
                    addi    r7 doForTwo:lo
                    bflag   r7 GT                   ;if r6 a[i] >= v || i == R
                    bflag   r7 EQ
                    cmp     r2 r5
                    bflag   r7 EQ
                    ldi     r7 doForOne:hi
                    shl     r7 r7 8
                    addi    r7 doForOne:lo
                    jump    r7


doForTwo:           subi    r3 one                  ;j--
                    ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    ldw     r6 r7 r3                ;R6 = a[j]
                    cmp     r6 r0                   ;if a[j] <= v || j == 0
                    ldi     r7 breakForJ--:hi
                    shl     r7 r7 8
                    addi    r7 breakForJ--:lo
                    bflag   r7 EQ
                    bflag   r7 LT
                    cmpi    r3 zero
                    bflag   r7 EQ
                    ldi     r7 doForTwo:hi
                    shl     r7 r7 8
                    addi    r7 doForTwo:lo
                    jump    r7

                    ; end for two
breakForJ--:        ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    ldw     r1 r7 r2                ;t = a[i]
                    ldw     r6 r7 r3                ;r6 = a[j]
                    stw     r6 r7 r2                ;a[i] = a[j]
                    stw     r1 r7 r3                ;a[j] = t  

                    ldi     r7 endForLoops:hi
                    shl     r7 r7 8
                    addi    r7 endForLoops:lo
                    cmp     r3 r2                   ;j <= i  break
                    bflag   r7 EQ
                    bflag   r7 LT
                    ldi     r7 doForOne:hi
                    shl     r7 r7 8
                    addi    r7 doForOne:lo
                    jump    r7

 


endForLoops:        ldi     r7 buffer:hi
                    shl     r7 r7 8
                    addi    r7 buffer:lo
                    ldw     r6 r7 r2
                    stw     r6 r7 r3
                    ldw     r6 r7 r5
                    stw     r6 r7 r2
                    stw     r1 r7 r5
                    ;i is in R2

                    MVSR    r0
                    ldi     r1 four
                    stw     r2 r0 r1                ;save i on in stack frame

                    ldi     r7 retPartition:hi
                    shl     r7 r7 8
                    addi    r7 retPartition:lo
 
                    jump    r7
                   



;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;
; golang quick sort algorithm used to model this assembly PROGRAM
;
;
;   package main
;
;   var a = []int{'A', 'E', 'D', 'B', 'C'}
;
;   func main() {
;   	qs(0, len(a)-1)
;   	fmt.Println(a)
;   }
;
;   func qs(l int, r int) {
;   	var i int
;   	fmt.Printf("QS %d %d\n", l, r)
;   	if r > l {
;   		i = partition(l, r)
;   		qs(l, i-1)
;   		qs(i+1, r)
;   	}
;   }
;
;   func partition(l int, r int) int {
;
;   	var v int
;   	var t int
;   	var i int
;   	var j int
;
;   	v = a[r]
;   	i = l - 1
;   	j = r
;   	for {
;   		for {
;   			i++
;   			if a[i] >= v || i == r {
;   				break
;   			}
;   		}
;   		for {
;   			j--
;   			if a[j] <= v || j == 0 {
;   				break
;   			}
;   		}
;   		t = a[i]
;   		a[i] = a[j]
;   		a[j] = t
;   		if j <= i {
;   			break
;   		}
;   	}
;   	a[j] = a[i]
;   	a[i] = a[r]
;   	a[r] = t
;   	return i
;   }
;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;


