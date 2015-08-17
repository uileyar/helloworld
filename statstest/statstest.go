package main

import (
	//"github.com/facebookgo/ensure"
	"fmt"
	"github.com/facebookgo/stats"
)

// Ensure calling End works even when a BumpTimeHook isn't provided.
func TestHookClientBumpTime() {
	(&stats.HookClient{}).BumpTime("foo").End()
}

func TestPrefixHookClient() {
	const (
		prefix1      = "prefix1-"
		prefix2      = "prefix2-"
		avgKey       = "avg-john"
		avgVal       = float64(1)
		sumKey       = "sum-john"
		sumVal       = float64(2)
		histogramKey = "histogram-john"
		histogramVal = float64(3)
		timeKey      = "time-john"
	)

	var keys []string
	hc := &stats.HookClient{
		BumpAvgHook: func(key string, val float64) {
			keys = append(keys, key)
			fmt.Printf("BumpAvgHook:key = %v, val = %v, avgVal = %v\n", key, val, avgVal)
			//ensure.DeepEqual(t, val, avgVal)
		},
		BumpSumHook: func(key string, val float64) {
			keys = append(keys, key)
			//ensure.DeepEqual(t, val, sumVal)
			fmt.Printf("BumpSumHook:key = %v, val = %v, sumVal = %v\n", key, val, sumVal)
		},
		BumpHistogramHook: func(key string, val float64) {
			keys = append(keys, key)
			//ensure.DeepEqual(t, val, histogramVal)
			fmt.Printf("BumpHistogramHook:key = %v, val = %v, histogramVal = %v\n", key, val, histogramVal)
		},

		BumpTimeHook: func(key string) interface {
			End()
		} {
			return multiEnderTest{
				EndHook: func() {
					keys = append(keys, key)
					fmt.Printf("EndHook:key = %v\n", key)
				},
			}
		},
	}

	pc := stats.PrefixClient([]string{prefix1, prefix2}, hc)
	pc.BumpAvg(avgKey, avgVal)
	pc.BumpSum(sumKey, sumVal)
	pc.BumpHistogram(histogramKey, histogramVal)
	pc.BumpTime(timeKey).End()

	fmt.Printf("%#v\n", keys)
}

func TestPrefixClient() {
	const (
		prefix1      = "prefix1-"
		prefix2      = "prefix2-"
		avgKey       = "avg-john"
		avgVal       = float64(1)
		sumKey       = "sum-john"
		sumVal       = float64(2)
		histogramKey = "histogram-john"
		histogramVal = float64(3)
		timeKey      = "time-john"
	)

	var keys []string
	hc := &stats.Client{
		BumpAvg: func(key string, val float64) {
			keys = append(keys, key)
			fmt.Printf("BumpAvg:key = %v, val = %v, avgVal = %v\n", key, val, avgVal)
			//ensure.DeepEqual(t, val, avgVal)
		},
		BumpSum: func(key string, val float64) {
			keys = append(keys, key)
			//ensure.DeepEqual(t, val, sumVal)
			fmt.Printf("BumpSum:key = %v, val = %v, sumVal = %v\n", key, val, sumVal)
		},
		BumpHistogram: func(key string, val float64) {
			keys = append(keys, key)
			//ensure.DeepEqual(t, val, histogramVal)
			fmt.Printf("BumpHistogram:key = %v, val = %v, histogramVal = %v\n", key, val, histogramVal)
		},

		BumpTime: func(key string) interface {
			End()
		} {
			return multiEnderTest{
				EndHook: func() {
					keys = append(keys, key)
					fmt.Printf("EndHook:key = %v\n", key)
				},
			}
		},
	}

	pc := stats.PrefixClient([]string{prefix1, prefix2}, hc)
	pc.BumpAvg(avgKey, avgVal)
	pc.BumpSum(sumKey, sumVal)
	pc.BumpHistogram(histogramKey, histogramVal)
	pc.BumpTime(timeKey).End()

	fmt.Printf("%#v\n", keys)
}

type multiEnderTest struct {
	EndHook func()
}

func (e multiEnderTest) End() {
	fmt.Printf("End: e = %#v\n", e)
	e.EndHook()
}

func main() {
	TestPrefixClient()
	//TestHookClientBumpTime()
}
