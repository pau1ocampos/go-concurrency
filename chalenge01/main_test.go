package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	sdtOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = sdtOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but is not there")
	}
}

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("xpto", &wg)
	wg.Wait()

	if msg != "xpto" {
		t.Errorf("Expecter xpto, but received %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	msg = "Hello, world!"

	sdtOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	printMessage()
	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = sdtOut

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find 'Hello, world!', but didn't")
	}
}

func Test_challenge1(t *testing.T) {
	sdtOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	challenge1()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = sdtOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but didn't")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but didn't")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but didn't")
	}
}
