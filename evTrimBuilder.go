package main

type evTrimBuilder struct {
	wheelType  string
	bodyType   string
	engineType string
}

func newEVTrimBuilder() *evTrimBuilder {
	return &evTrimBuilder{}
}

func (b *evTrimBuilder) setWheelType() {
	b.wheelType = "16 inch All Season"
}

func (b *evTrimBuilder) setBodyStyle() {
	b.bodyType = "3 Door Hatchback"
}

func (b *evTrimBuilder) setEngineType() {
	b.engineType = "Electric 265 bhp"
}

func (b *evTrimBuilder) getCarModel() model {
	return model{
		wheelType:  b.wheelType,
		bodyType:   b.bodyType,
		engineType: b.engineType,
	}
}
