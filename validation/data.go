package validation

import "github.com/gookit/validate"

type Data struct {
	data validate.DataFace
}

func (d *Data) AddError(key string, rule string, message string) {
	//TODO implement me
	panic("implement me")
}

func NewData(data validate.DataFace) *Data {
	return &Data{data}
}

func (d *Data) Get(key string) (val any, exist bool) {
	return d.data.Get(key)
}

func (d *Data) Set(key string, val any) error {
	_, err := d.data.Set(key, val)

	return err
}
