package database

type DbOptions struct {
	filePath string
}

func NewDbOptions() *DbOptions {
	o := &DbOptions{}
	return o
}

func (o *DbOptions) SetFilePath(filePath string) *DbOptions {
	o.filePath = filePath
	return o
}
