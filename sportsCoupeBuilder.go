package main

type sportsCoupeBuilder struct {
	wheelType  string
	bodyType   string
	engineType string
}

func newSportsCoupeBuilder() *sportsCoupeBuilder {
	return &sportsCoupeBuilder{}
}

func (b *sportsCoupeBuilder) setWheelType() {
	b.wheelType = "19 inch Performance"
}

func (b *sportsCoupeBuilder) setBodyStyle() {
	b.bodyType = "2 Door Coupe"
}

func (b *sportsCoupeBuilder) setEngineType() {
	b.engineType = "V8 380 bhp"
}

func (b *sportsCoupeBuilder) getCarModel() model {
	return model{
		wheelType:  b.wheelType,
		bodyType:   b.bodyType,
		engineType: b.engineType,
	}
}
