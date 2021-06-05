package main

import (
	"math/rand"
	"strings"
)

func RandomString(len int) string {
	b := strings.Builder{}
	for i := 0; i < len; i++ {
		b.WriteRune(RandomRune())
	}
	return b.String()
}

func RandomRune() rune {
	const (
		min = int('0')
		max = int('z')
	)
	return rune(rand.Intn(max-min) + min)
}
