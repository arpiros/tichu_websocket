package controllers

import (
	"testing"
	"math/rand"
)

func TestShuffleArray(t *testing.T) {
	list := rand.Perm(25)
	for key, value := range list {
		println(key, value)
	}
}
