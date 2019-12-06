package units

const (
	Celcius = "C"
	Pascal  = "Pa"
	Meter   = "m"
)

// Prop contains the numerical value for a property of a physical state
type Prop struct {
	value float64
	unit  string
}
