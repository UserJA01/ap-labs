package main

import "fmt"
import "os"
import "strconv"

func calc(operator int, values []int) int {
    var option int
    option = operator
    
	
	var result int

	result = values[2]

    for x := 3; x < len(values); x++ {

        if option == 0 {

        	result += values[x]
        }else if option == 1 {

            result -= values[x]
        }else if option == 2 {

        	result *= values[x]
        }
    }

    return result
}

func main() {
    var length int
    length = len(os.Args)
    

    num := make([]int, length)
    for x := 2; x < length; x++ {
        temp, err := strconv.Atoi(os.Args[x])
        if err == nil {
            num[x] = temp
	    }
    }
    
    var option int
    if os.Args[1] == "add" {
        option = 0;
    }else if os.Args[1] == "sub" {
        option = 1;
    }else if os.Args[1] == "mult" {
        option = 2;
    } else{
        fmt.Println("Please write down your operator(add,sub or multi) followed by the numbers you want\n");
        fmt.Println("Example:\n");
        fmt.Println("add 5 2 1 5\n");
        os.Exit(0)
    }
    var result int
    result = 0
    if len(os.Args) >= 2 {
        result = calc(option, num)
    }
    
    for x := 2; x < length; x++ {
        //fmt.Printf(num[x])
        fmt.Printf("%v",num[x])
        
        if(x < length - 1){
            if(option == 0){
                fmt.Printf(" + ")
            }else if(option == 1){
                fmt.Printf(" - ")
            }else if(option == 2){
                fmt.Printf(" * ")
            }
        }else{
            fmt.Printf(" = ")
        }
    }
    fmt.Printf(" %v \n", result)
}
