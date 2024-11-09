package elon

import "fmt"

// Drive the car for the given distance
func (c *Car) Drive() {
	if c.speed > c.battery {
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// DisplayDistance returns the distance the car has driven
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	return float32(c.battery) >= float32(trackDistance)/float32(c.speed)*float32(c.batteryDrain)
}
