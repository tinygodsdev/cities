package cities

import (
	"testing"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
)

func TestNewCity(t *testing.T) {
	attributes := map[string]data.Attribute{
		AttributeTemperature: {Values: []string{"20.5", "21.5", "19.0"}},
		AttributeHumidity:    {Values: []string{"50", "55", "53"}},
		AttributePressure:    {Values: []string{"1013", "1012", "1015"}},
		AttributeDescription: {Values: []string{"Sunny"}},
		// Add other attributes here as needed
	}

	timestamp := time.Now()
	city := NewCity("TestCity", attributes, timestamp)

	// Check the city name
	if city.Name != "TestCity" {
		t.Errorf("Expected city name 'TestCity', got '%s'", city.Name)
	}

	// Check the city timestamp
	if !city.Timestamp.Equal(timestamp) {
		t.Errorf("Expected timestamp '%v', got '%v'", timestamp, city.Timestamp)
	}

	// Check the calculated average temperature
	expectedTemperature := "20.33"
	if city.Temperature != expectedTemperature {
		t.Errorf("Expected temperature '%s', got '%s'", expectedTemperature, city.Temperature)
	}

	// Check the calculated average humidity
	expectedHumidity := "52.67"
	if city.Humidity != expectedHumidity {
		t.Errorf("Expected humidity '%s', got '%s'", expectedHumidity, city.Humidity)
	}

	// Check the calculated average pressure
	expectedPressure := "1013.33"
	if city.Pressure != expectedPressure {
		t.Errorf("Expected pressure '%s', got '%s'", expectedPressure, city.Pressure)
	}

	// Check the description
	expectedDescription := "Sunny"
	if city.Description != expectedDescription {
		t.Errorf("Expected description '%s', got '%s'", expectedDescription, city.Description)
	}

	// Add more checks for other attributes as needed
}
