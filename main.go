package main

import (
	"fmt"
)

const (
	mm = 0.001
	cm = 0.01
	m  = 1
	km = 1000
	in = 0.0254
	ft = 0.3048
	yd = 0.9144
	mi = 1609.34
)

const (
	g  = 1
	kg = 1000
	lb = 453.592
	oz = 28.3495
)

var LengthUnit = map[string]float64{
	"mm": mm,
	"cm": cm,
	"m":  m,
	"km": km,
	"in": in,
	"ft": ft,
	"yd": yd,
	"mi": mi,
}

var WeightUnit = map[string]float64{
	"g":  g,
	"kg": kg,
	"lb": lb,
	"oz": oz,
}

func convertLength(value float64, from float64, to float64) float64 {
	return value * from / to
}

func convertWeight(value float64, from float64, to float64) float64 {
	return value * from / to
}

// Temperature conversion function
func convertTemperature(value float64, from string, to string) float64 {
	switch from {
	case "C":
		switch to {
		case "C":
			return value
		case "F":
			return value*9/5 + 32
		case "K":
			return value + 273.15
		}
	case "F":
		switch to {
		case "C":
			return (value - 32) * 5 / 9
		case "F":
			return value
		case "K":
			return (value-32)*5/9 + 273.15
		}
	case "K":
		switch to {
		case "C":
			return value - 273.15
		case "F":
			return (value-273.15)*9/5 + 32
		case "K":
			return value
		}
	}
	return 0
}

func main() {
	var choice int
	fmt.Println("Enter 1 for Length Conversion")
	fmt.Println("Enter 2 for Weight Conversion")
	fmt.Println("Enter 3 for Temperature Conversion")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var value float64
		fmt.Println("Enter the value to convert:")
		fmt.Scan(&value)
		var from, to string
		fmt.Println("Enter the unit to convert from (mm, cm, m, km, in, ft, yd, mi):")
		fmt.Scan(&from)
		fmt.Println("Enter the unit to convert to (mm, cm, m, km, in, ft, yd, mi):")
		fmt.Scan(&to)
		fmt.Printf("%.4f %s = %.4f %s\n", value, from, convertLength(value, LengthUnit[from], LengthUnit[to]), to)
	case 2:
		var value float64
		fmt.Println("Enter the value to convert:")
		fmt.Scan(&value)
		var from, to string
		fmt.Println("Enter the unit to convert from (g, kg, lb, oz):")
		fmt.Scan(&from)
		fmt.Println("Enter the unit to convert to (g, kg, lb, oz):")
		fmt.Scan(&to)
		fmt.Printf("%.4f %s = %.4f %s\n", value, from, convertWeight(value, WeightUnit[from], WeightUnit[to]), to)
	case 3:
		var value float64
		fmt.Println("Enter the value to convert:")
		fmt.Scan(&value)
		var from, to string
		fmt.Println("Enter the unit to convert from (C, F, K):")
		fmt.Scan(&from)
		fmt.Println("Enter the unit to convert to (C, F, K):")
		fmt.Scan(&to)
		fmt.Printf("%.2f %s = %.2f %s\n", value, from, convertTemperature(value, from, to), to)
	default:
		fmt.Println("Invalid choice")
	}
}
