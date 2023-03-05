package domain

type Member struct {
	Id   int
	Name string
	Age  int
}

type MemberRepository interface {
	FetchAll() ([]Member, error)
	Add(mem Member) error
	CheckIfExists(mem Member) (bool, error)
}

func NewMember(name string, age int) *Member {
	return &Member{Name: name, Age: age}
}
