package id

import "math/rand"

type ID int64

func New() ID {
	return ID(rand.Int63())
}
