        @i
        M=1 // i=1
        @sum
        M=0
(LOOP)
        @i
        D=M // D=i
        @100
        D=D-A // D=i-100
        @END
        D;JGT // もし(i-100)>0ならばENDに移動
        @i
        D=M // D=i
        @sum
        M=D+M // sum=sum+i
        @i
        M=M+1 // i=i+1
        @LOOP
        0;JMP // LOOPに移動
(END)
        @END
        0;JMP // 無限ループ
