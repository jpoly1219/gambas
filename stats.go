package gambas

import (
	"fmt"
	"math"
	"sort"
)

type StatsFunc func(dataset []interface{}) StatsResult

type StatsResult struct {
	UsedFunc string
	Result   float64
	Err      error
}

// Count counts the number of non-NA elements in a column.
func Count(dataset []interface{}) StatsResult {
	count := 0
	for _, v := range dataset {
		if v != nil || v != math.NaN() {
			count++
		}
	}

	return StatsResult{"Count", float64(count), nil}
}

// Mean returns the mean of the elements in a column.
func Mean(dataset []interface{}) StatsResult {
	mean := 0.0

	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Mean", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Mean", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	for _, v := range data {
		mean += v
	}

	mean /= float64(len(data))
	roundedMean := math.Round(mean*1000) / 1000

	return StatsResult{"Mean", roundedMean, nil}
}

// Median returns the median of the elements in a column.
func Median(dataset []interface{}) StatsResult {
	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Median", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Median", math.NaN(), fmt.Errorf("no elements in this column")}
	}
	if total%2 == 0 {
		lower := data[total/2-1]
		upper := data[total/2]

		median := (lower + upper) / 2
		roundedMedian := math.Round(median*1000) / 1000

		return StatsResult{"Median", roundedMedian, nil}
	} else {
		median := data[(total+1)/2-1]
		roundedMedian := math.Round(median*1000) / 1000

		return StatsResult{"Median", roundedMedian, nil}
	}
}

// Std returns the sample standard deviation of the elements in a column.
func Std(dataset []interface{}) StatsResult {
	std := 0.0
	meanResult := Mean(dataset) // this also checks that all data can be converted to float64.
	if meanResult.Err != nil {
		return StatsResult{"Std", math.NaN(), meanResult.Err}
	}

	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Std", math.NaN(), err}
	}
	sort.Float64s(data)

	numerator := 0.0
	for _, v := range data {
		temp := math.Pow(v-meanResult.Result, 2)
		numerator += temp
	}
	std = math.Sqrt(numerator / float64(len(data)-1))
	roundedStd := math.Round(std*1000) / 1000

	return StatsResult{"Std", roundedStd, nil}
}

// Min returns the smallest element in a column.
func Min(dataset []interface{}) StatsResult {
	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Min", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Min", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	return StatsResult{"Min", data[0], nil}
}

// Max returns the largest element is a column.
func Max(dataset []interface{}) StatsResult {
	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Max", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Max", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	return StatsResult{"Max", data[total-1], nil}
}

// Q1 returns the lower quartile (25%) of the elements in a column.
// This does not include the median during calculation.
func Q1(dataset []interface{}) StatsResult {
	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Q1", math.NaN(), err}
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		lower := data[:len(data)/2]
		q1, err := median(lower)
		if err != nil {
			return StatsResult{"Q1", math.NaN(), err}
		}
		return StatsResult{"Q1", q1, nil}
	} else {
		lower := data[:(len(data)-1)/2]
		q1, err := median(lower)
		if err != nil {
			return StatsResult{"Q1", math.NaN(), err}
		}
		return StatsResult{"Q1", q1, nil}
	}
}

// Q2 returns the middle quartile (50%) of the elements in a column.
// This accomplishes the same thing as s.Median().
func Q2(dataset []interface{}) StatsResult {
	q2Result := Median(dataset)
	if q2Result.Err != nil {
		return StatsResult{"Q2", math.NaN(), q2Result.Err}
	}

	return StatsResult{"Q2", q2Result.Result, nil}
}

// Q3 returns the upper quartile (75%) of the elements in a column.
// This does not include the median during calculation.
func Q3(dataset []interface{}) StatsResult {
	data, err := interface2F64Slice(dataset)
	if err != nil {
		return StatsResult{"Q3", math.NaN(), err}
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		upper := data[len(data)/2:]
		q3, err := median(upper)
		if err != nil {
			return StatsResult{"Q3", math.NaN(), err}
		}
		return StatsResult{"Q3", q3, nil}
	} else {
		upper := data[(len(data)+1)/2:]
		q3, err := median(upper)
		if err != nil {
			return StatsResult{"Q3", math.NaN(), err}
		}
		return StatsResult{"Q3", q3, nil}
	}
}
