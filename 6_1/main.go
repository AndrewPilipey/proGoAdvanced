type Quidditch struct {
	RedTeam  Team
	BlueTeam Team
}

func NewQuidditch() Quidditch {
	return Quidditch{
		RedTeam:  NewTeam(),
		BlueTeam: NewTeam(),
	}
}

func (q Quidditch) ShowScore() {
	q.RedTeam.PrintInfo()
	q.BlueTeam.PrintInfo()
}

func (q Quidditch) RedSnitch() {
	RedTeam += 150
	fmt.Printf("Red Team Won!")
}

func (q Quidditch) BlueSnitch() {
	BlueTeam += 150
	fmt.Printf("Blue Team Won!")
}