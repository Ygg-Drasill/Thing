package penalties

type Penalty struct {
	Name  string
	Value int
}

var Penalties = []Penalty{
	{"§1 Gasbøde", 2},
	{"§2 Dummebøde", 2},
	{"§3 Spildebøde", 2},
	{"§4 Push to main", 25},
	{"§5 Glemme bøde", 5},
	{"§6 Fremmøde med tømmermænd", 5},
	{"§7 Dårligt menneske bøde", 2},
	{"§8 Mangel af kilder", 2},
	{"§9 AI", 20},
	{"§10 Bailer bøde", 20},
}

func FindPenalty(name string) (Penalty, bool) {
	for _, penalty := range Penalties {
		if penalty.Name == name {
			return penalty, true
		}
	}
	return Penalty{}, false
}
