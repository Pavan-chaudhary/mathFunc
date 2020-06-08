package metrics

import (
    "fmt"
    "math"
)

func getMean(arr *[]float64) float64{
    var res float64 = 0
    for _,val := range *arr {
        res += val
    }
    return res/float64(len(*arr))
}

func getVariance(arr *[]float64) float64{
    mean := getMean(arr)
    var sq_sum float64 = 0
    for _,val := range *arr {
        sq_sum += math.Pow(val-mean,2)
    }
    return sq_sum
}

func getStandardDeviation(arr *[]float64) float64{
    return math.Sqrt(getVariance(arr))
}

