package latex

// Cv represents the entire LaTeX resume
type Cv struct {
	Header                  Cvheader
	ProfessionalSummary     ProfessionalSummary `json:"professional_summary"`
	ProfessionalExperiences []Cventry           `json:"professional_experiences"`
	Skills                  []Cvitem
	Certifications          []Cvitem
	Eduction                []Cvitem
	Languages               []Cvitem
	Interests               []Cvitem
}

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

// ProfessionalSummary represents the summary at the top of the CV
type ProfessionalSummary string

// Define a \cvitem

// Cvitem represents a cvitem in the modercv LaTeX package
type Cvitem struct {
	Title   string // The title of the cvitem
	Content string // the content of the cvitem
}

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

// List defines a LaTeX list
type List []string
