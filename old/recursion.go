package main

import "fmt"

func main() {
	fmt.Println(getSum(5))
}

func getSum(n int) int { //getSum(5)=getSum(4)+5
	if n == 1 { //getSum(4)=getSum(3)+4
		return 1 //getSum(3)=getSum(2)+3
	} //getSum(2)=getSum(1)+2
	return getSum(n-1) + n //getSum(1)=1
}
