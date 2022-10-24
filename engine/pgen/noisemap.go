package pgen

import (
	"math"

	"github.com/ojrac/opensimplex-go"
)

type Octave struct {
	Freq, Scale float64
}

type NoiseMap struct {
	seed     int64
	noise    opensimplex.Noise
	ocataves []Octave
	exponent float64
}

// TODO: ensure that sum of all octave amplitudes equals 1!
func NewNoiseMap(seed int64, ocataves []Octave, exponent float64) *NoiseMap {
	return &NoiseMap{
		seed:     seed,
		noise:    opensimplex.NewNormalized(seed),
		ocataves: ocataves,
		exponent: exponent,
	}
}

func (n *NoiseMap) Get(x, y int) float64 {
	ret := 0.0
	for i := range n.ocataves {
		xNoise := n.ocataves[i].Freq * float64(x)
		yNoise := n.ocataves[i].Freq * float64(y)
		ret += n.ocataves[i].Scale * n.noise.Eval2(xNoise, yNoise)

	}

	ret = math.Pow(ret, n.exponent)
	return ret
}
