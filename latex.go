package latex

import "fmt"

// Printlist prints a LaTeX list (\begin{itemize} ... \end{itemize})
func Printlist(list []string) {
	fmt.Println("\\begin{itemize}")
	for _, item := range list {
		fmt.Println("\\item", item)
	}
	fmt.Println("\\end{itemize}")
}
