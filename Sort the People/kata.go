package Sort_the_People

import "sort"

type Persons struct {
	Names   []string
	Heights []int
}

func (p Persons) Len() int           { return len(p.Names) }
func (p Persons) Less(i, j int) bool { return p.Heights[i] > p.Heights[j] }
func (p Persons) Swap(i, j int) {
	p.Heights[i], p.Heights[j] = p.Heights[j], p.Heights[i]
	p.Names[i], p.Names[j] = p.Names[j], p.Names[i]
}

func sortPeople(names []string, heights []int) []string {
	p := &Persons{Names: names, Heights: heights}
	sort.Sort(p)
	return p.Names
}
