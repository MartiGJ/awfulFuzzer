package main

import (
	"math/rand"
	"time"
)

var mutation, baseTest []byte

func generator(inputTestCase []byte, sigWork chan<- []byte) {
	mutators := []func(byte) byte{rndMagic, rndByte, func(b byte) byte { return b }}
	mutation = make([]byte, len(inputTestCase))
	baseTest = make([]byte, len(inputTestCase))
	copy(baseTest, inputTestCase)
	rand.Seed(time.Now().UnixNano())
	count := 0
	for {
		count++
		copy(mutation, baseTest)
		//Mutate test case
		for i := range mutation {
			rnd := rand.Intn(len(mutators))
			mutation[i] = mutators[rnd](mutation[i])
		}

		//Signal worker
		testCase := make([]byte, len(mutation))
		copy(testCase, mutation)
		sigWork <- testCase

	}
}

func rndMagic(b byte) byte {
	mgk := []byte{'b', 'W', 'F', 'y', 'd'}
	return mgk[rand.Intn(len(mgk))]
}

func rndByte(b byte) byte {
	b = byte(rand.Int63())
	for b == 0 {
		b = byte(rand.Int63())
	}
	return b
}
