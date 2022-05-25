package main

import (
	BO "mypkg/basicOperation"
	CT "mypkg/convTemp"
	"strconv"
	"testing"
)

func TestCase1(t *testing.T){
	want := "55"
	expected := BO.Add(5, 50)
	if want != expected {
		t.Error("La suma no retorno el resultado esperado")
	}
}
func TestCase2(t *testing.T){
	want := "3"
	expected := BO.Substract(5,2)
	if want != expected {
		t.Error("La resta no retorno el resultado esperado")
	}
}
func TestCase3(t *testing.T){
	want := "8"
	expected := BO.MultiplyBy(4, 2)
	if want != expected {
		t.Error("La multiplicacion no retorno el resultado esperado")
	}
}
func TestCase4(t *testing.T){
	want := "10"
	expected := BO.DivideBy(20,2)
	if want != expected {
		t.Error("La division no retorno el resultado esperado")
	}
}
func TestCase5(t *testing.T){
	want := "Error, no puedes dividir entre 0"
	expected := BO.DivideBy(30,0)
	if want != expected {
		t.Error("La division no retorno el resultado esperado")
	}
}
func TestCase6(t *testing.T){
	want := 50
	expected := CT.ToFarenheit(10, CT.Celsius)
	if float64(want) != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Farenheit")
	}
}
func TestCase7(t *testing.T){
	want := 283.15
	expected := CT.ToKelvin(10, CT.Celsius)
	if want != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Kelvin")
	}
}
func TestCase8(t *testing.T){
	want := 10
	expected := CT.ToCelsius(10, CT.Celsius)
	if float64(want) != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Celsius")
	}
}
func TestCase9(t *testing.T){
	want := 50
	expected := CT.ToFarenheit(50, CT.Fahrenheit)
	if float64(want) != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Farenheit")
	}
}
func TestCase10(t *testing.T){
	want := 283.15
	expected := CT.ToKelvin(50, CT.Fahrenheit)
	if want != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Kelvin")
	}
}
func TestCase11(t *testing.T){
	want := 10
	expected := CT.ToCelsius(50, CT.Fahrenheit)
	if float64(want) != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Celsius")
	}
}
func TestCase12(t *testing.T){
	want := "-369.67"
	expected := CT.ToFarenheit(50, CT.Kelvin)
	if want != strconv.FormatFloat(expected , 'f', -1, 32) {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Farenheit")
	}
}
func TestCase13(t *testing.T){
	want := 50
	expected := CT.ToKelvin(50, CT.Kelvin)
	if float64(want) != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Kelvin")
	}
}
func TestCase14(t *testing.T){
	want := "-223.15"
	expected := CT.ToCelsius(50, CT.Kelvin)
	if want != strconv.FormatFloat(expected , 'f', -1, 32) {
		t.Error("No retorno el resultado esperado en cuanto a la conversion a Celsius")
	}
}
func TestCase15(t *testing.T){
	want := "283.15K-50F"
	expected := CT.ToString(10, CT.Celsius)
	if want != expected {
		t.Error("No retorno el resultado esperado en cuanto a la conversion string")
	}
}
