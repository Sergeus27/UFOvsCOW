package main

import "fmt"

applyOperation(a, b int, operation func(int, int) int) int{
	return operation(a,b)
}

func main() {
result:=applyOperation(10,5,func(a,b){
	return(a-b)
})	
fmt.Println(result)

}
