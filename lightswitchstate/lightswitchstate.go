package main

import "fmt"

var lightSwitchIsOn bool = false

func main() {
  printLightSwitchState()
  toggleLightSwithcState()
  printLightSwitchState()
  toggleLightSwithcState()
  printLightSwitchState()
}

func printLightSwitchState() {
  fmt.Println("The light switch is off:", lightSwitchIsOn)
}

func toggleLightSwithcState() {
  lightSwitchIsOn = !lightSwitchIsOn
}
