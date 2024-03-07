package main

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(color Color) Builder
	Wheels(wheels Wheels) Builder
	TopSpeed(speed Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func main() {
}
