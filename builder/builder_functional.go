package builder

import "fmt"

type Person2 struct {
	name, position string
}

type personMod func(*Person2)

type Person2Builder struct {
	actions []personMod
}

func (b *Person2Builder) Called(name string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) {
		p.name = name
	})
	return b
}

func (b *Person2Builder) Build() *Person2 {
	p := Person2{}

	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *Person2Builder) WorksAsA(position string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) {
		p.position = position
	})

	return b
}

func RunFunctional() {
	b := Person2Builder{}
	p := b.Called("Dmitri").WorksAsA("developer").Build()
	fmt.Println(*p)
}
