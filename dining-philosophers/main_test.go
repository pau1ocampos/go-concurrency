package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect lenght of slice, expceted 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryindDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, tst := range theTests {
		orderFinished = []string{}
		eatTime = tst.delay
		thinkTime = tst.delay
		sleepTime = tst.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect lenght of slice, expceted 5 but got %d", tst.name, len(orderFinished))
		}

	}
}
