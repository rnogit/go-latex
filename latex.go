package latex

// Cv represent the entire cv
type Cv struct {
	Header struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Address   string `json:"address"`
		Zipcode   string `json:"zipcode"`
		City      string `json:"city"`
		State     string `json:"state"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Homepage  string `json:"homepage"`
		Quote     string `json:"quote"`
	} `json:"header"`
	ProfessionalSummarySection struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"professional_summary_section"`
	ProfessionalExperiencesSection struct {
		Title       string `json:"title"`
		Experiences []struct {
			Name         string   `json:"name"`
			Date         string   `json:"date"`
			Position     string   `json:"position"`
			Company      string   `json:"company"`
			State        string   `json:"state"`
			City         string   `json:"city"`
			Description  string   `json:"description"`
			Achievements []string `json:"achievements"`
		} `json:"experiences"`
	} `json:"professional_experiences_section"`
	SkillsSection struct {
		Title  string `json:"title"`
		Skills []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		} `json:"skills"`
	} `json:"skills_section"`
	CertificationSection struct {
		Title            string `json:"title"`
		VerificationLink string `json:"verification_link"`
		Certifications   []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		} `json:"certifications"`
	} `json:"certification_section"`
	EductionSection struct {
		Title     string `json:"title"`
		Education []struct {
			Date        string `json:"date"`
			College     string `json:"college"`
			State       string `json:"state"`
			DegreeLevel string `json:"degree_level"`
			DegreeName  string `json:"degree_name"`
		} `json:"education"`
	} `json:"eduction_section"`
	LanguagesSection struct {
		Title     string `json:"title"`
		Languages []struct {
			Language    string `json:"language"`
			Proficiency string `json:"proficiency"`
		} `json:"languages"`
	} `json:"languages_section"`
	InterestsSection struct {
		Title     string `json:"title"`
		Interests []struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		} `json:"interests"`
	} `json:"interests_section"`
}
