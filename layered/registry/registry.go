package registry

type Registry struct {
	db DB
}

func NewRegistry() Registry {
	return Registry{
		db: NewDB(),
	}
}

func (r Registry) DB() DB {
	return r.db
}
