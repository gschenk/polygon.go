package geobjects

import "time"

const t50 = 1586910095 * 1e9 // approx 50 years in nanoseconds

// generate a unique id integer (creation time in nanoseconds)
func id() int64 { return time.Now().UnixNano() - t50 }
