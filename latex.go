package latex

import "fmt"

// Printlist prints a LaTeX list (\begin{itemize} ... \end{itemize})
func Printlist(list []string) {
	fmt.Println("\\begin{itemize}")
	for _, item := range list {
		fmt.Println("	\\item", item)
	}
	fmt.Println("\\end{itemize}")
}

// Cvitem represents a cvitem in the modercv LaTeX template
type Cvitem struct {
	title   string
	content string
}

// Printcvitem prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func Printcvitem(cvitem Cvitem) {
	fmt.Printf("\\cvitem{%s}{\textnormal{%s}}\n", cvitem.title, cvitem.content)
}
