// Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// @Author Pankaj

// The @Haversine formula calculates the distance between two points,
// using latitude and logitude on a sphere

package main

import (
	"fmt"
	"math"
)

const (
	// assigning π value
	π = math.Pi

	// R defines radius of Earth
	R = 6371e3
)

// Point defines position of an object on the sphere
// with the healp of latitude and logitude
type Point struct {
	Latitude  float64 // Latitude of point on a sphere
	Longitude float64 // Longitude of point on a sphere
}

// Harversine defines the Harversine formula
type Harversine struct{}

func degreeToRadius(θ float64) float64 {
	return θ * π / 180 // convert degree to radius and return the value
}

// CalculateDistance takes two points as parameter,
// performs Harversine farmula to calulate the distanc
// return distance between points in meter
func (h *Harversine) CalculateDistance(from, to *Point) float64 {

	φ1 := degreeToRadius(from.Latitude)  // convert Latitude degree to radius, we need this radius in the formula
	λ1 := degreeToRadius(from.Longitude) // convert Longitude degree to radius
	φ2 := degreeToRadius(to.Latitude)    // convert Latitude degree to radius
	λ2 := degreeToRadius(to.Longitude)   // convert Longitude degree to radius
	Δφ := φ2 - φ1                        // Get difference between Latitude
	Δλ := λ2 - λ1                        // Get difference between Longitude

	// Apply Harversine formula
	// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
	// c = 2 ⋅ atan2( √a, √(1−a) )
	// d = R ⋅ c
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c // return distance in meter
}

func main() {

	point1 := &Point{40.689202777778, -74.044219444444}
	point2 := &Point{38.889069444444, -77.034502777778}

	harversine := &Harversine{}

	fmt.Println(harversine.CalculateDistance(point1, point2))

}
