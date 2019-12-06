package units

const (
	KelvinOffset = 273.15
)

// ToKelvin converts the property to Kelvin
func (p *Prop) ToKelvin() {
	switch p.unit {
	case Celcius:
		p.value += KelvinOffset
	}
}

// ToFahrenheit converts the property to Fahrenheit
func (p *Prop) ToFahrenheit() {
	switch p.unit {
	case Celcius:
		p.value = p.value*(9/5) + 32
	}
}
