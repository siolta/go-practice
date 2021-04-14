// get the number of seconds since unix epoch

package main

import (
	"fmt"
	"time"
)

func main() {
	// get elapsed time since unixk epoch
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// to get milliseconds you need to manually divide
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// can convert int secs since epoch into corresponding time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
