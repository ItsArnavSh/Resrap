package main

import (
	"math/rand"
)

type PRNG struct {
	seed   uint64
	number uint64
}

func (p *PRNG) setSeed(seed uint64) {
	p.number = seed
	p.seed = seed
}
func (p *PRNG) generateSeed() {
	seed := rand.Uint64()
	p.number = seed
	p.seed = seed
}
func (p *PRNG) nextPRN() uint64 {
	//Using the XOR shift method for PRN generation
	p.number ^= p.number << 13
	p.number ^= p.number >> 7
	p.number ^= p.number << 17
	return p.number
}

// random returns a float64 in [0,1)
func (p *PRNG) Random() float64 {
	// Take the next 53 random bits (same precision as math/rand.Float64)
	v := p.nextPRN() >> 11        // keep top 53 bits
	return float64(v) / (1 << 53) // normalize to [0,1)
}

// randomInt returns an int in [min, max)
func (p *PRNG) RandomInt(min, max int) int {
	if max <= min {
		return min // avoid division by zero or negative range
	}
	r := p.Random()
	return min + int(r*float64(max-min))
}

func newPRNG(seed uint64) PRNG {
	prng := PRNG{}
	if seed == 0 {
		prng.generateSeed()
	} else {
		prng.setSeed(seed)
	}
	return prng
}
