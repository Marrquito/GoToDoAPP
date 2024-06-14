package utils

type Output_str struct {
	Output string
}

func (o *Output_str) Update(v string) {
	o.Output = v
}