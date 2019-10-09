package solver

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

// Solver is a pow solver
type Solver struct {
	SolverConfig
}

// SolverConfig is config struct for Solver
type SolverConfig struct {
	Prefix     string
	Difficulty int
}

// New returns a POW solver
func New(c SolverConfig) (*Solver, error) {
	if c.Difficulty <= 0 {
		return nil, errors.New("Difficulty should > 0")
	}
	return &Solver{
		SolverConfig{
			Prefix:     c.Prefix,
			Difficulty: c.Difficulty,
		},
	}, nil
}

// Solve returns two fields,
// @ret1 the number of iteration after solving the pow.
// @ret2 binary form of the answer.
func (s *Solver) Solve() (attempts int, binaryString string) {
	var i int
	var ss string
	n := time.Now()
	for {
		if e := time.Since(n); e > 1*time.Second {
			break
		}

		sum := sha256.Sum256([]byte(fmt.Sprintf("%s%v", s.Prefix, i)))
		ss = toBinString(sum)
		if isValid(ss, s.Difficulty) {
			break
		}
		i++
	}
	return i, ss
}

func toBinString(s [sha256.Size]byte) string {
	var ss string
	for _, b := range s {
		ss = fmt.Sprintf("%s%08b", ss, b)
	}
	return ss
}

func isValid(ss string, d int) bool {
	for i := range ss {
		if i >= d {
			break
		}
		if ss[i] != '0' {
			return false
		}
	}
	return true
}
