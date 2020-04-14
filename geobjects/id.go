package geobjects

import "time"

const tzero = 1577837 * 1e12 // approx unix time at 2020-01-01 in [ns]

// generate a unique id integer (creation time in nanoseconds)
func id() int64 { return time.Now().UnixNano() - tzero }
