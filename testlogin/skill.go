package testlogin

import "fmt"

type Person3 struct {
	Name  string
	Sex   string
	Skill map[string]int
}

func NewPerson(name string, sex string, skill map[string]int) Person3 {
	return Person3{Name: name, Sex: sex, Skill: skill}
}

func (p *Person3) NewSkill() {

}

func (p *Person3) upLevel(skill string) {

}

type PersonSkill map[string]map[string]int

func (ps *PersonSkill) getSkillLevel(name, skill string) int {
	return 0
}

func Skill() {
	f := "football"
	b := "boxing"
	yong := NewPerson(
		"yong",
		"MALE",
		map[string]int{
			f: 8,
			b: 10,
		},
	)

	pak := NewPerson(
		"pak",
		"MALE",
		map[string]int{
			f: 6,
			b: 6,
		},
	)

	personSkill := NewPersonSkill([]Person3{yong, pak})

	personSkill.getSkillLevel(pak.Name, "football")
	fmt.Println()

}

func NewPersonSkill(persons []Person3) PersonSkill {
	personslist := make(map[string]map[string]int)
	// ทำเอง
	return personslist
}
