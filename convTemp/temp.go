package temp

import "strconv"

type TemperatureScale string

const (
	Celsius    TemperatureScale = "Celsius"
	Fahrenheit TemperatureScale = "Fahrenheit"
	Kelvin     TemperatureScale = "Kelvin"
)

func ToFarenheit(temp float64, T TemperatureScale) float64 {
	var fahrenheitTemp float64
	switch T {
	case Celsius:
		fahrenheitTemp = (1.8 * temp) + 32
	case Fahrenheit:
		fahrenheitTemp = temp
	case Kelvin:
		fahrenheitTemp = 1.8*(temp-273.15) + 32
	}
	return fahrenheitTemp
}
func ToCelsius(temp float64, T TemperatureScale) float64 {
	var CelsiusTemp float64
	switch T {
	case Celsius:
		CelsiusTemp = temp
	case Fahrenheit:
		CelsiusTemp =  (temp - 32)/1.8
	case Kelvin:
		CelsiusTemp = temp - 273.15
	}
	return CelsiusTemp
}
func ToKelvin(temp float64, T TemperatureScale) float64 {
	var KelvinTemp float64
	switch T {
	case Celsius:
		KelvinTemp = temp + 273.15
	case Fahrenheit:
		KelvinTemp = (5*(temp - 32))/9 + 273.15
	case Kelvin:
		KelvinTemp = temp
	}
	return KelvinTemp
}
func ToString(temp float64, T TemperatureScale)string {
	var Resultado string
	switch T {
	case Celsius:
		KelvinTemp := ToKelvin(temp, T)
		KelvinTString := strconv.FormatFloat(KelvinTemp, 'f', -1, 32) + "K"
		FarenheitTemp := ToFarenheit(temp, T)
		FahrenheitTString := strconv.FormatFloat(FarenheitTemp, 'f', -1, 32) +"F"

		Resultado = KelvinTString + "-" + FahrenheitTString
	case Fahrenheit:
		CelsiusTemp := ToCelsius(temp, T)
		CelsiusTString := strconv.FormatFloat(CelsiusTemp,'f',-1,32 ) +"C"
		KelvinTemp := ToKelvin(temp, T)
		KelvinTString := strconv.FormatFloat(KelvinTemp, 'f', -1, 32) +"K"

		Resultado = CelsiusTString + "-" + KelvinTString
	case Kelvin:
		CelsiusTemp := ToCelsius(temp, T)
		CelsiusTString := strconv.FormatFloat(CelsiusTemp,'f',-1,32 ) +"C"
		FarenheitTemp := ToFarenheit(temp, T)
		FahrenheitTString := strconv.FormatFloat(FarenheitTemp, 'f', -1, 32) +"F"

		Resultado = CelsiusTString + "-" + FahrenheitTString 
	default:
		return "Temparatura desconocida"
	}
	return Resultado
	

}