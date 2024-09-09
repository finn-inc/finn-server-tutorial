package registry

type Registry struct {
	DB DB
}

func NewRegistry() Registry {
	return Registry{
		DB: NewDB(),
	}
}
