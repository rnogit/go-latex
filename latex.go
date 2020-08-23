package latex

import "fmt"

func list(list []int) {
	fmt.Println("\\begin{itemize}")
	for _, item := range list {
		fmt.Println("\\item", item)
	}
	fmt.Println("\\end{itemize}")
}
