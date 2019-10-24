#include <stdio.h>
#include <stdlib.h>

//clang -g testPrgrm.c -fsanitize=address -fsanitize-coverage=bb,trace-pc-guard -o testPrgm
int main(int argc, char const *argv[])
{
    
    if(argc < 2){
        printf("Run %s with arguments\n",argv[0]);
        return 0;
    }
    
    if (argv[1][0]=='b'){
        if (argv[1][1]=='W'){
            if (argv[1][2]=='F'){
                if (argv[1][3]=='y'){
                    if (argv[1][4]=='d'){
                        abort();
                    }
                }
            }
        }
    }
    return 0;
}
