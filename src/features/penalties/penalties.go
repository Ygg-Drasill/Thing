package penalties

type Penalty struct {
	Name  string
	Value int
	Info  string
}

var Penalties = []Penalty{
	{"§1 Gasbøde", 2, PenaltyDescriptions["§1 Gasbøde"]},
	{"§2 Dummebøde", 2, PenaltyDescriptions["§2 Dummebøde"]},
	{"§3 Spildebøde", 2, PenaltyDescriptions["§3 Spildebøde"]},
	{"§4 Push to main", 2, PenaltyDescriptions["§4 Push to main"]},
	{"§5 Glemme bøde", 2, PenaltyDescriptions["§5 Glemme bøde"]},
	{"§6 Fremmøde med tømmermænd", 2, PenaltyDescriptions["§6 Fremmøde med tømmermænd"]},
	{"§7 Dårligt menneske bøde", 2, PenaltyDescriptions["§7 Dårligt menneske bøde"]},
	{"§8 Mangel af kilder", 2, PenaltyDescriptions["§8 Mangel af kilder"]},
	{"§9 AI", 2, PenaltyDescriptions["§9 AI"]},
	{"§10 Bailer bøde", 2, PenaltyDescriptions["§10 Bailer bøde"]},
}

func FindPenalty(name string) (Penalty, bool) {
	for _, penalty := range Penalties {
		if penalty.Name == name {
			return penalty, true
		}
	}
	return Penalty{}, false
}
