package controllers

import (
	"testing"
	"math/rand"
)

func TestShuffleArray(t *testing.T) {
	list := rand.Perm(25)
	for i, _ := range list {
		list[i]++
	}

	for key, value := range list {
		println(key, value)
	}
}