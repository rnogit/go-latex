package latex

import "fmt"

type Cv struct {
	Header                  Cvheader
	ProfessionalExperiences []Cventry
}

// List defines a LaTeX list
type List []string

// Print prints a LaTeX list (\begin{itemize} ... \end{itemize})
func (l List) Print() {
	fmt.Println("\\begin{itemize}")
	for _, item := range l {
		fmt.Println("	\\item", item)
	}
	fmt.Println("\\end{itemize}")
}

// Define a \cvitem

// Cvitem represents a cvitem in the modercv LaTeX package
type Cvitem struct {
	Title   string // The title of the cvitem
	Content string // the content of the cvitem
}

// Print prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (i Cvitem) Print() {
	fmt.Printf("\\cvitem{%s}{\\textnormal{%s}}\n", i.Title, i.Content)
}

// Define a \cventry

// Cventry represents a cventry in the modercv LaTeX package
type Cventry struct {
	Date         string // The date
	Position     string // the title or name of the position
	Company      string // The company name
	City         string // The city name
	State        string // State
	Description  string // Description of the position
	Achievements List   // A list of achievements
}

// Print prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (e Cventry) Print() {
	fmt.Printf("\\cventry{%s}{%s}{\\textsc{%s}}{%s, %s}{}{\\textnormal{%s}}{\n", e.Date, e.Position, e.Company, e.City, e.State, e.Description)
	e.Achievements.Print()
	fmt.Println("}")
}

// Define the header of the cv

// Cvheader represents all the attributs of the modercv package used in the header of the resume
type Cvheader struct {
	Firstname string
	Lastname  string
	Address   string
	Zipcode   string
	City      string
	State     string
	Country   string
	Phone     string
	Email     string
	Homepage  string
	Quote     string
}

// Print prints a LaTeX moderncv cvitem (\cvitem{title}{\textnormal{content}})
func (h Cvheader) Print() {
	fmt.Printf("\\firstname{%s}\n", h.Firstname)
	fmt.Printf("\\familyname{%s}\n", h.Lastname)
	fmt.Printf("\\address{%s}{%s %s, %s}\n", h.Address, h.Zipcode, h.City, h.Country)
	fmt.Printf("\\mobile{%s}\n", h.Phone)
	fmt.Printf("\\email{%s}\n", h.Email)
	fmt.Printf("\\homepage{%s}{}\n", h.Homepage)
	fmt.Printf("\\quote{%s}\n", h.Quote)
}
