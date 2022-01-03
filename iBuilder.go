package main

type iBuilder interface {
	setWheelType()
	setBodyStyle()
	setEngineType()
	getCarModel() model
}

func getBuilder(builderType string) iBuilder {
	if builderType == "basicTrim" {
		return &basicTrimBuilder{}
	}

	if builderType == "evTrim" {
		return &evTrimBuilder{}
	}

	if builderType == "sportsCoupeTrim" {
		return &sportsCoupeBuilder{}
	}

	return nil
}
