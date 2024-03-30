package data

type TemplateData struct {
	Title          string
	Header         string
	Items          []string
	Person         []string
	SelectedPerson string
	Penalty        map[string]int
	Penalties      map[string]int
}

func NewTemplateData(person []string, selectedPersonStr string, penalty map[string]int) TemplateData {
	return TemplateData{
		Title:          "Thing",
		Header:         "Welcome to Thing!",
		Items:          []string{"You", "Will", "Pay"},
		Person:         person,
		SelectedPerson: selectedPersonStr,
		Penalty:        penalty,
		Penalties:      make(map[string]int),
	}
}
