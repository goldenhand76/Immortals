package database

type DbOptions struct {
	filePath string
}

func NewDbOptions() *DbOptions {
	o := &DbOptions{
		filePath: "internal\\database\\sqlite\\immo.db",
	}
	return o
}

func (o *DbOptions) SetFilePath(filePath string) *DbOptions {
	o.filePath = filePath
	return o
}
