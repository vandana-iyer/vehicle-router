package main

import (
	"bytes"
	. "com.vorto.vehiclerouter/pkg/models"
	"os"
	"testing"
)

func TestProcess(t *testing.T) {
	// Redirect standard output to capture the printed solution
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Process(getMockLoads())

	// Read the captured output
	err := w.Close()
	if err != nil {
		return
	}
	var buf bytes.Buffer
	buf.ReadFrom(r)

	os.Stdout = old
	got := buf.String()

	// Expected output
	want := "[3,1,4]\n[2]\n"

	if got != want {
		t.Errorf("Got:\n%s \nBut want:\n%s", got, want)
	}
}

func getMockLoads() []Load {
	return []Load{
		{LoadId: 1, Pickup: Location{-50.1, 80.0}, Dropoff: Location{90.1, 12.2}},
		{LoadId: 2, Pickup: Location{-24.5, -19.2}, Dropoff: Location{98.5, 1.8}},
		{LoadId: 3, Pickup: Location{0.3, 8.9}, Dropoff: Location{40.9, 55.0}},
		{LoadId: 4, Pickup: Location{5.3, -61.1}, Dropoff: Location{77.8, -5.4}},
	}
}
