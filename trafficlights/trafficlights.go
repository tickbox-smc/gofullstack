package main

import (
	"fmt"
)

const (
	_ = iota
	TrafficLightStateRedLight
	TrafficLightStateGreenLight
	TrafficLightStateYellowLight
)

func main() {

  fmt.Println("Red light State Code: ", TrafficLightStateRedLight)
  fmt.Println("Green light State Code", TrafficLightStateGreenLight)
  fmt.Println("Yellow light State Code", TrafficLightStateYellowLight)

}
