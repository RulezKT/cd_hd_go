package main

import (
	"fmt"

	"github.com/RulezKT/cd_consts_go"
	"github.com/RulezKT/cd_de_440s_go"
	"github.com/RulezKT/cd_hd_go"
)

func main() {

	bsp := cd_de_440s_go.Load440s()

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

}
