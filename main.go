package main

import "fmt"

func main() {
	basicTrim := getBuilder("basicTrim")
	evTrim := getBuilder("evTrim")
	sportsCoupe := getBuilder("sportsCoupeTrim")

	// Build a Basic Trim Level Model
	director := newDirector(basicTrim)
	basicTrimModel := director.buildCarModel()
	fmt.Println("Basic Trim Level")
	fmt.Printf("Body Style: %s\n", basicTrimModel.bodyType)
	fmt.Printf("Wheel Type: %s\n", basicTrimModel.wheelType)
	fmt.Printf("Engine: %s\n", basicTrimModel.engineType)
	fmt.Println()

	// Build a EV (electric) Trim Level Model
	director.setBuilder(evTrim)
	evTrimModel := director.buildCarModel()

	fmt.Println("EV Trim Level")
	fmt.Printf("Body Style: %s\n", evTrimModel.bodyType)
	fmt.Printf("Wheel Type: %s\n", evTrimModel.wheelType)
	fmt.Printf("Engine: %s\n", evTrimModel.engineType)
	fmt.Println()

	// Build a Sports Coupe Trim Level Model
	director.setBuilder(sportsCoupe)
	sportsCoupeModel := director.buildCarModel()

	fmt.Println("Sports Coupe Trim Level")
	fmt.Printf("Body Style: %s\n", sportsCoupeModel.bodyType)
	fmt.Printf("Wheel Type: %s\n", sportsCoupeModel.wheelType)
	fmt.Printf("Engine: %s\n", sportsCoupeModel.engineType)
	fmt.Println()

}
