package metrics

import math

func GetSquareArea(side float64) area float64 {
    area := side * side 
    return
}

func GetRectangleArea(length float64, breadth float64) area float64 {
    area := length * breadth
    return
} 

func GetCircleArea(radius float64) area float64 {
    area := radius * math.Pi * math.Pi
    return
}
