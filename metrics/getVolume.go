package metrics

import "math"

func GetSphereVolume(radius float64) (vol float64) {
    vol = (4 * math.Pi * radius * radius * radius)/3
    return
}

func GetCylinderVolume(radius float64, height float64) (vol float64){
    vol = 2 * math.Pi* radius * height
    return
}
