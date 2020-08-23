package latex

import "fmt"

type list []string

// NewList create a new LaTeX list (\begin{itemize} ... \end{itemize})
func NewList(items ...string) []string {
	i := items
	return i
}

// Printlist prints a LaTeX list (\begin{itemize} ... \end{itemize})
func (l list) Print() {
	fmt.Println("\\begin{itemize}")
	for _, item := range l {
		fmt.Println("	\\item", item)
	}
	fmt.Println("\\end{itemize}")
}

// Cvitem represents a cvitem in the modercv LaTeX template
type cvitem struct {
	title   string // The title of the cvitem
	content string // the content of the cvitem
}

// NewCvitem create a new cvitem
func NewCvitem(title string, content string) cvitem {
	i := cvitem{title, content}
	return i
}

// Printcvitem prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (cvitem cvitem) Print() {
	fmt.Printf("\\cvitem{%s}{\\textnormal{%s}}\n", cvitem.title, cvitem.content)
}
