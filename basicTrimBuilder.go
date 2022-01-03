package main

type basicTrimBuilder struct {
	wheelType  string
	bodyType   string
	engineType string
}

func newBasicTrimBuilder() *basicTrimBuilder {
	return &basicTrimBuilder{}
}

func (b *basicTrimBuilder) setWheelType() {
	b.wheelType = "17 inch All Season"
}

func (b *basicTrimBuilder) setBodyStyle() {
	b.bodyType = "4 Door Sedan"
}

func (b *basicTrimBuilder) setEngineType() {
	b.engineType = "V6 230 bhp"
}

func (b *basicTrimBuilder) getCarModel() model {
	return model{
		wheelType:  b.wheelType,
		bodyType:   b.bodyType,
		engineType: b.engineType,
	}
}
