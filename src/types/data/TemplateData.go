package data

import "github.com/Ygg-Drasill/Thing/src/features/penalties"

type TemplateData struct {
	Title           string
	Header          string
	Items           []string
	Person          []string
	SelectedPerson  string
	Penalties       []penalties.Penalty
	SelectedPenalty penalties.Penalty
	Owes            map[string]int
}

func NewTemplateData(person []string, selectedPersonStr string, owes map[string]int, selectedPenalty penalties.Penalty) TemplateData {
	return TemplateData{
		Title:           "Thing",
		Header:          "Welcome to Thing!",
		Items:           []string{"You", "Will", "Pay"},
		Person:          person,
		SelectedPerson:  selectedPersonStr,
		Penalties:       penalties.Penalties,
		SelectedPenalty: selectedPenalty,
		Owes:            owes,
	}
}
