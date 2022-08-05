package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3)
	b := make([]int, 3)
	var al int
	var ba int
	// comp := make([]int, 2)
	for i := 0; i < len(a); i++{
		fmt.Scan(&a[i])
		
	}
	for i := 0; i < len(a); i++{
		fmt.Scan(&b[i])
		
	}
	for k := 0; k < len(a); k++{
		if a[k] > b[k]{
			al++
		}else if a[k] == b[k]{
			continue
		}else{
			ba++
		}
	}


	fmt.Println(al,ba )
	


}
