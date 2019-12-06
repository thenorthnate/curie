package therm

import (
	"github.com/thenorthnate/curie"
)

// State represents a thermodynamic state and includes most computable thermodynamic attributes
type State struct {
	A curie.Prop // Cross-sectional area
	D curie.Prop // Diameter (m)
	P curie.Prop // pressure (Pa)
	T curie.Prop // temperature (C)
	V curie.Prop // volume (m^3)
}

type Compressor struct {
	Inlet  State
	Outlet State
}
