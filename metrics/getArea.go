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

func GetRightTriangleArea(base float64, height float64) area float64 {
    area := (base * height)/2
    return
}

func GetEquilateralTriangleArea(side float64) area float64 {
    area := (math.Sqrt(3) * side * side)/4
    return
}
