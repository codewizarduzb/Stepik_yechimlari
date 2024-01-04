package main

// вам это понадобится

import (
	"fmt"
	"time"
)

const now = 1589570165

func main_3_7_4() {

	var daqiqa, soniya int64
	fmt.Scanf("%d мин. %d сек.", &daqiqa, &soniya)

	unixTime := time.Unix(
		now+60*daqiqa+soniya,
		0,
	)

	fmt.Println(unixTime.Format(time.UnixDate))
}
