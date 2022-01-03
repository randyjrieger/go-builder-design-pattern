package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildCarModel() model {
	d.builder.setWheelType()
	d.builder.setBodyStyle()
	d.builder.setEngineType()
	return d.builder.getCarModel()
}
