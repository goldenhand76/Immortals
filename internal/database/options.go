package database

type DbOptions struct {
	name string
}

func NewDbOptions() *DbOptions {
	o := &DbOptions{}
	return o
}

func (o *DbOptions) SetName(name string) *DbOptions {
	o.name = "reza"
	return o
}
