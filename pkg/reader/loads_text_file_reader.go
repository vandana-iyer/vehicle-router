package reader

import (
	"bufio"
	"com.vorto.vehiclerouter/pkg/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	header       = "loadNumber pickup dropoff"
	totalColumns = 3
)

// loadsTextFileReader is an implementation of LoadsFileReader,
// specifically designed to read loads data from text files
type loadsTextFileReader struct{}

// NewLoadsFileReader is a factory function to create a new instance of loadsTextFileReader
func NewLoadsFileReader() LoadsFileReader {
	return &loadsTextFileReader{}
}

func (l *loadsTextFileReader) Read(filePath string) ([]models.Load, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseLoadsFromScanner(bufio.NewScanner(file))
}

func parseLoadsFromScanner(scanner *bufio.Scanner) ([]models.Load, error) {
	var loads []models.Load

	// Check if the header is as expected
	if scanner.Scan() && scanner.Text() != header {
		return nil, fmt.Errorf("header mismatch")
	}

	for scanner.Scan() {
		load, err := processLine(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error processing line: %v", err)
		}
		loads = append(loads, load)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return loads, nil
}

func processLine(line string) (models.Load, error) {
	fields := strings.Split(line, " ")
	if len(fields) < totalColumns {
		return models.Load{}, fmt.Errorf("invalid line format: %s", line)
	}

	const loadIdIndex = 0
	id, err := strconv.Atoi(fields[loadIdIndex])
	if err != nil {
		return models.Load{}, fmt.Errorf("error parsing load id: %v", err)
	}

	const pickUpIndex = 1
	pickup, err := parseLocation(fields[pickUpIndex])
	if err != nil {
		return models.Load{}, fmt.Errorf("error parsing pickup coordinates: %v", err)
	}

	const dropOffIndex = 2
	dropoff, err := parseLocation(fields[dropOffIndex])
	if err != nil {
		return models.Load{}, fmt.Errorf("error parsing dropoff coordinates: %v", err)
	}

	return models.Load{
		LoadId:  id,
		Pickup:  pickup,
		Dropoff: dropoff,
	}, nil
}

func parseLocation(latLng string) (models.Location, error) {
	latLng = strings.Trim(latLng, "()")
	locationCoordinates := strings.Split(latLng, ",")

	if len(locationCoordinates) != 2 {
		return models.Location{}, fmt.Errorf("invalid coordinate format: %s", latLng)
	}

	const latitudeIndex = 0
	latitude, err := getCoordinate(locationCoordinates, latitudeIndex)
	if err != nil {
		return models.Location{}, fmt.Errorf("error parsing latitude: %v", err)
	}

	const longitudeIndex = 1
	longitude, err := getCoordinate(locationCoordinates, longitudeIndex)
	if err != nil {
		return models.Location{}, fmt.Errorf("error parsing longitude: %v", err)
	}

	return models.Location{Latitude: latitude, Longitude: longitude}, nil
}

func getCoordinate(locationCoordinates []string, index int) (float64, error) {
	latitude, err := strconv.ParseFloat(strings.TrimSpace(locationCoordinates[index]), 64)
	return latitude, err
}
