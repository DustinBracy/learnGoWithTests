package specifications

type GreetAdapter func(name string) string

func (a GreetAdapter) Greet(name string) (string, error) {
	return a(name), nil
}

type CurseAdapter func(name string) string

func (a CurseAdapter) Curse(name string) (string, error) {
	return a(name), nil
}
