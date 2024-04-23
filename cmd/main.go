package main

import (
	"fmt"
	"math"

	"github.com/RulezKT/cd_consts_go"
	"github.com/RulezKT/cd_de_440s_go"
	"github.com/RulezKT/cd_hd_go"
)

func test_calc_design_time(bsp cd_consts_go.BspFile) {
	var max_diff int64 = 0
	var min_diff int64 = math.MaxInt64
	var max_i int64
	var min_i int64
	var v1 int64
	fmt.Println("Start testing...")
	for i := int64(-4734072000.0 + 8_000_000); i < 4735368000-8_000_000; i += 10_000_000 {

		v1, _ = cd_hd_go.CalcDesignTimeV3(i, bsp)
		// v2, _ = cd_hd_go.CalcDesignTimeV2(i, bsp)
		// diff := math.Abs(float64(v2 - v1))
		// if diff > 1000 {
		// 	panic(fmt.Sprintf("V1: %v V2: %v  difference: %v ", v1, v2, diff))
		// }

		diff := int64(math.Abs(float64(i - v1)))

		if diff > 8_000_000 {

			fmt.Println("difference: ", diff)
			fmt.Println("i: ", i)
			panic(fmt.Sprintf("V1: %v  difference: %v ", v1, diff))
		}

		if diff > max_diff {
			max_diff = diff
			max_i = i
		}

		if diff < min_diff {
			min_diff = diff
			min_i = i
		}

	}

	fmt.Println("Max diff: ", max_diff)
	fmt.Println("Min diff: ", min_diff)

	fmt.Println("Max diff i: ", max_i)
	fmt.Println("Min diff i: ", min_i)

}

func main() {

	bsp := cd_de_440s_go.Load440s()
	// test_calc_design_time(bsp)

	var info cd_consts_go.CdInfo
	info.HdInfo.Init()

	utc_time := cd_consts_go.GregDate{
		Year:    1978,
		Month:   05,
		Day:     17,
		Hour:    12,
		Minutes: 45,
		Seconds: 01,
	}
	var date_to_calc = cd_consts_go.TimeData{
		LocalTime: cd_consts_go.GregDate{}, // LocalTime GregDate //для design всегда 0
		UtcTime:   utc_time,                // UtcTime   cd_consts_go.GregDate

		TypeOfTyme:    1,  // TypeOfTyme    int    //Изначальный источник данных 2 - local time, 1- UTC Time,  0 - Ephemeries time
		Offset:        0,  // Offset        int    //смещение локального времени от UTC в секундах
		SecFromJd2000: 0,  // SecFromJd2000 int64  // Ephemeries time
		Place:         "", // Place         string // не пустой, только если время изначально Local, для design всегда пустой
	}

	//-682470731
	cd_hd_go.CalcCosmo(date_to_calc, bsp, &info)

	fmt.Println(date_to_calc)

	// my seconds -682470731

	v2, _ := cd_hd_go.CalcDesignTimeV3(
		-4719200000, bsp)

	fmt.Println(v2)

	v2, _ = cd_hd_go.CalcDesignTimeV3(-4_719_196_000, bsp)
	fmt.Println(v2)

	v2, _ = cd_hd_go.CalcDesignTimeV3(1_682_470_731, bsp)
	fmt.Println(v2)

}

// -4734072000.0
// 4735368000.0

// V1: -690181870.2032913
// V2: -690181872.162277

// V1: 674521431.0567086
// V2: 674521428.6707723

// V1: 1674845033.4967082

//Max diff:  7_951_038
//Min diff:  7_482_038
