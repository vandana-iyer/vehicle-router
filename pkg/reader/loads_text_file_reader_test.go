package reader

import (
	"com.vorto.vehiclerouter/pkg/models"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	// Setup a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_loads_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal("Error removing temp file")
		}
	}(tempFile.Name())

	content := getContent()
	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Create a new LoadFileReader
	loadReader := NewLoadsFileReader()

	// Use the LoadFileReader to read from the temp file
	got, err := loadReader.Read(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read from temp file: %v", err)
	}

	// Validate the results
	want := []models.Load{
		{LoadId: 1, Pickup: models.Location{Latitude: -50.1, Longitude: 80.0}, Dropoff: models.Location{Latitude: 90.1, Longitude: 12.2}},
		{LoadId: 2, Pickup: models.Location{Latitude: -24.5, Longitude: -19.2}, Dropoff: models.Location{Latitude: 98.5, Longitude: 1.8}},
		{LoadId: 3, Pickup: models.Location{Latitude: 0.3, Longitude: 8.9}, Dropoff: models.Location{Latitude: 40.9, Longitude: 55.0}},
	}

	if len(got) != len(want) {
		t.Fatalf("got %d loads, but want %d", len(got), len(want))
	}

	for i, item := range got {
		if item != want[i] {
			t.Errorf("got %v, but want %v", item, want[i])
		}
	}
}

func getContent() string {
	content :=
		`loadNumber pickup dropoff
1 (-50.1,80.0) (90.1,12.2) 
2 (-24.5,-19.2) (98.5,1.8) 
3 (0.3,8.9) (40.9,55.0)`
	return content
}
