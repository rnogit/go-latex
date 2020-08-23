package latex

import "fmt"

// Define a LaTeX list
type list []string

// NewList creates a new LaTeX list (\begin{itemize} ... \end{itemize})
func NewList(items ...string) list {
	list := items
	return list
}

// Print prints a LaTeX list (\begin{itemize} ... \end{itemize})
func (l list) Print() {
	fmt.Println("\\begin{itemize}")
	for _, item := range l {
		fmt.Println("	\\item", item)
	}
	fmt.Println("\\end{itemize}")
}

// Define a \cvitem

// cvitem represents a cvitem in the modercv LaTeX package
type cvitem struct {
	title   string // The title of the cvitem
	content string // the content of the cvitem
}

// NewCvitem create a new cvitem
func NewCvitem(title string, content string) cvitem {
	i := cvitem{title, content}
	return i
}

// Print prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (cvitem cvitem) Print() {
	fmt.Printf("\\cvitem{%s}{\\textnormal{%s}}\n", cvitem.title, cvitem.content)
}

// Define a \cventry

// cventry represents a cventry in the modercv LaTeX package
type cventry struct {
	date         string // The date
	position     string // the title or name of the position
	company      string // The company name
	city         string // The city name
	state        string // State
	description  string // Description of the position
	achievements list   // A list of achievements
}

// NewCventry create a new cvitem
func NewCventry(date string, position string, company string, city string, state string, description string, achievements list) cventry {
	e := cventry{date, position, company, city, state, description, achievements}
	return e
}

// Printcvitem prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (e cventry) Print() {
	fmt.Printf("\\cventry{%s}{%s}{\\textsc{%s}}{%s, %s}{}{\\textnormal{%s}}{\n", e.date, e.position, e.company, e.city, e.state, e.description)
	e.achievements.Print()
	fmt.Println("}")
}
