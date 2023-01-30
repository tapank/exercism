package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	if t == 0 {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	}
	return ""
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (sp Speed) String() string {
	return fmt.Sprintf("%d %s", sp.magnitude, sp.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (md MeteorologyData) String() string {
	template := "%s: %s, Wind %s at %s, %d%% Humidity"
	l := md.location
	t := md.temperature
	d := md.windDirection
	s := md.windSpeed
	h := md.humidity
	return fmt.Sprintf(template, l, t, d, s, h)
}
