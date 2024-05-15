package utils

type Output struct {
	Output string
}

func (o *Output) Update(v string) {
	o.Output = v
}