package main

import "fmt"

type PersonF struct {
	name, position string
}

type personMod func(*PersonF)
type PersonFBuilder struct {
	actions []personMod
}

func (b *PersonFBuilder) Called(name string) *PersonFBuilder {
	b.actions = append(b.actions, func(p *PersonF) {
		p.name = name
	})
	return b
}

func (b *PersonFBuilder) Build() *PersonF {
	p := PersonF{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *PersonFBuilder) WorksAsA(position string) *PersonFBuilder {
	b.actions = append(b.actions, func(p *PersonF) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonFBuilder{}
	p := b.Called("Dmitri").WorksAsA("dev").Build()
	fmt.Println(*p)
}
