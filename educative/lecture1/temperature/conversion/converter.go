package conversion

/*
	Temperature in ˚F = (Temp in °C * 9/5) + 32
*/

// Celcius value in float32
type Celcius float32

// Fahrenheit value in float32
type Fahrenheit float32

// ToFahrenheit convert from celcius to fahrenheit
func ToFahrenheit(celcius Celcius) Fahrenheit {
	temp := ((celcius * 9) / 5) + 32
	return Fahrenheit(temp)
}
