#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int main(int argc, char **argv)
{
    int option = 0;
    if(argc >= 2){
        if(strcmp(argv[1], "add") == 0){
            //printf("CORRECT!\n");
            option = 0;
        }else if(strcmp(argv[1], "sub") == 0){
            option = 1;
        }else if(strcmp(argv[1], "mult") == 0){
            option = 2;
        }
        else{
            //printf("Please write down your input in the next format: ");
            printf("Please write down your operator(add,sub or multi) followed by the numbers you want\n");
            printf("Example:\n");
            printf("add 5 2 1 5\n");
            exit(EXIT_FAILURE);
        }
        
        int result = atoi(argv[2]);
        //printf("Result: %d \n", result);
        for(int x = 3; x < argc; x++){
            if(isdigit(atoi(argv[x])) != 0){
                /*int ax = atoi(argv[x]);
                result = 123;*/
                exit(EXIT_FAILURE);
            }
            if(option == 0){
                result += atoi(argv[x]);
            }else if(option == 1){
                result -= atoi(argv[x]);
            }else if(option == 2){
                result *= atoi(argv[x]);
            }
            
        }

        printf("%d \n", result);
    }
    return 0;
}
