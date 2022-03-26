package idgen

import (
	"math/rand"
	"time"

	"go-ddd-blog/internal/domain/idgen"
)

type randomIDGen struct {
}

func NewRandomIDGen() idgen.IDGenerator {
	return &randomIDGen{}
}

func (g *randomIDGen) Generate() (int64, error) {
	ts := time.Now().Unix()
	return ts - ts%100 + int64(rand.Intn(100)), nil
}
