package main

import "fmt"

type animal interface {
	Run() string
	Shout() string
}

type cat struct {
	speed int
	sound string
}

func (c *cat) Run() string {
	return fmt.Sprintf("Cat run %d kilometer/hour.", c.speed)
}
func (c *cat) Shout() string {
	return fmt.Sprintf("Cat shout sound %s.", c.sound)
}
func main() {
	xiaobai := cat{
		speed: 10,
		sound: "miao~miao~miao",
	}
	fmt.Print(xiaobai.Run())
	fmt.Print(xiaobai.Shout())
}
